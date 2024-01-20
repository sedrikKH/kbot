/*
Copyright © 2023 Serhii Adamchuk adamchuk.serg@gmail.com
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"

	"github.com/hirosassa/zerodriver"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	// Teletoken bot
	TeleToken = os.Getenv("TELE_TOKEN")

	// MetricsHost exporter host:port
	MetricsHost = os.Getenv("METRICS_HOST")
)

// Currency struct to represent each currency in the JSON
type Currency struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	Cc           string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}

var (
	btnUSD = telebot.InlineButton{
		Unique: "usd_button",
		Text:   "USD",
		Data:   "kurs USD",
	}
	btnEUR = telebot.InlineButton{
		Unique: "eur_button",
		Text:   "EUR",
	}
	btnAUD = telebot.InlineButton{
		Unique: "aud_button",
		Text:   "AUD",
	}
	btnList = telebot.InlineButton{
		Unique: "list_button",
		Text:   "List",
	}
)

var buttons = [][]telebot.InlineButton{
	{btnUSD, btnEUR, btnAUD},
	{btnList},
}

var markup = telebot.ReplyMarkup{
	InlineKeyboard: buttons,
}

// Initialize OpenTelemetry
func initMetrics(ctx context.Context) {

	// Create a new OTLP Metric gRPC exporter with the specified endpoint and options
	exporter, _ := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(MetricsHost),
		otlpmetricgrpc.WithInsecure(),
	)

	// Define the resource with attributes that are common to all metrics.
	// labels/tags/resources that are common to all metrics.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(fmt.Sprintf("kbot_%s", appVersion)),
	)

	// Create a new MeterProvider with the specified resource and reader
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource),
		sdkmetric.WithReader(
			// collects and exports metric data every 10 seconds.
			sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(10*time.Second)),
		),
	)

	// Set the global MeterProvider to the newly created MeterProvider
	otel.SetMeterProvider(mp)

}

func pmetrics(ctx context.Context, payload string) {
	// Get the global MeterProvider and create a new Meter with the name "kbot_light_signal_counter"
	meter := otel.GetMeterProvider().Meter("kbot_currency_counter")

	// Get or create an Int64Counter instrument with the name "kbot_light_signal_<payload>"
	counter, _ := meter.Int64Counter(fmt.Sprintf("kbot_currency_%s", payload))

	// Add a value of 1 to the Int64Counter
	counter.Add(ctx, 1)
}

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := zerodriver.NewProductionLogger()

		// fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			logger.Fatal().Str("Error", err.Error()).Msg("Please check TELE_TOKEN")
			return
		} else {
			logger.Info().Str("Version", appVersion).Msg("kbot started")

		}

		kbot.Handle(&btnUSD, func(m telebot.Context) error {
			err = displayExchangeRate(m, "USD")
			if err != nil {
				return err
			}
			return err

		})

		kbot.Handle(&btnEUR, func(m telebot.Context) error {
			err = displayExchangeRate(m, "EUR")
			if err != nil {
				return err
			}
			return err

		})

		kbot.Handle(&btnAUD, func(m telebot.Context) error {
			err = displayExchangeRate(m, "AUD")
			if err != nil {
				return err
			}
			return err

		})

		kbot.Handle(&btnList, func(m telebot.Context) error {
			err = displayCurrencyList(m)
			if err != nil {
				return err
			}
			return err

		})

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			logger.Info().Str("Payload", m.Text()).Msg(m.Message().Payload)

			log.Print(m.Message().Payload, m.Text())
			//payload := m.Message().Payload
			payload := m.Text()
			pmetrics(context.Background(), payload)

			switch strings.ToLower(payload) {
			case "/start", "/hello":
				err = m.Send(fmt.Sprintf("Hello I'm kbot %s", appVersion), &markup)

			case "kurs":
				err = displayCurrencyList(m)
				if err != nil {
					return err
				}

			default:
				if strings.HasPrefix(strings.ToLower(payload), "kurs") {
					currencyCode := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(payload), "kurs"))
					if currencyCode != "" {
						err = displayExchangeRate(m, currencyCode)
						if err != nil {
							return err
						}
					} else {
						err := m.Send("Please provide a valid currency code after 'kurs', e.g., 'kurs USD'.", &markup)
						if err != nil {
							return err
						}
					}
				}
			}

			return err
		})

		kbot.Start()
	},
}

func init() {
	ctx := context.Background()
	initMetrics(ctx)
	rootCmd.AddCommand(kbotCmd)
}

func displayCurrencyList(m telebot.Context) error {
	res, err := http.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json")
	if err != nil {
		log.Fatalf("Error", err)
		return err
	}

	var currencies []Currency
	err = json.NewDecoder(res.Body).Decode(&currencies)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var currencyList string
	for _, currency := range currencies {
		currencyList += fmt.Sprintf("%s - %s\n", currency.Txt, currency.Cc)
	}

	err = m.Send("List of available currencies:\n"+currencyList, &markup)
	return err
}

func displayExchangeRate(m telebot.Context, currencyCode string) error {
	res, err := http.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json")
	if err != nil {
		log.Fatalf("Error", err)
		return err
	}

	var currencies []Currency
	err = json.NewDecoder(res.Body).Decode(&currencies)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// If a specific currency code is provided, filter the currencies
	var filteredCurrencies []Currency
	for _, c := range currencies {
		if strings.ToUpper(c.Cc) == strings.ToUpper(currencyCode) {
			filteredCurrencies = append(filteredCurrencies, c)
		}
	}
	currencies = filteredCurrencies

	if len(currencies) == 0 {
		err := m.Send(fmt.Sprintf("No exchange rate information found for currency code %s.", currencyCode), &markup)
		if err != nil {
			return err
		}
		return nil
	}

	for _, currency := range currencies {
		// Initialize the variables
		cc := currency.Cc
		exchangeDate := currency.ExchangeDate
		rate := currency.Rate
		err := m.Send(fmt.Sprintf("Валюта **%s** (Код валюти: %s) \nкурс обміну на %s: **%.2f** грн за 1 %s", currency.Txt, cc, exchangeDate, rate, cc), &markup)
		if err != nil {
			return err
		}
	}

	return nil
}
