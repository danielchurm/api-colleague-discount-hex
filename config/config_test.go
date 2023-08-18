package config_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/config"
)

var _ = Describe("Configuration", func() {

	Context("App config", func() {

		It("returns a pointer to an instance of AppConfig", func() {
			appConfig, err := config.NewAppConfig()

			Expect(err).ToNot(HaveOccurred())
			Expect(appConfig).ToNot(BeNil())
		})

		It("builds default port config", func() {
			appConfig, err := config.NewAppConfig()

			Expect(err).ToNot(HaveOccurred())
			Expect(appConfig.Port).ToNot(BeNil())

			Expect(appConfig.Port).To(Equal("8080"))

		})

		Context("Logger config", func() {
			It("builds default logger config", func() {
				appConfig, err := config.NewAppConfig()

				Expect(err).ToNot(HaveOccurred())
				Expect(appConfig.Logger).ToNot(BeNil())

				Expect(appConfig.Logger.LogLevel).To(Equal("info"))
			})

			It("reads logger level from environment vars", func() {
				currentLogLevel := os.Getenv("LOGS_APP_LEVEL")
				currentLogHttpBodies := os.Getenv("LOG_HTTP_BODIES")
				defer os.Setenv("LOGS_APP_LEVEL", currentLogLevel)
				defer os.Setenv("LOG_HTTP_BODIES", currentLogHttpBodies)

				os.Setenv("LOGS_APP_LEVEL", "some-log-level")
				os.Setenv("LOG_HTTP_BODIES", "true")

				appConfig, err := config.NewAppConfig()

				Expect(err).ToNot(HaveOccurred())
				Expect(appConfig.Logger).ToNot(BeNil())

				Expect(appConfig.Logger.LogLevel).To(Equal("some-log-level"))
				Expect(appConfig.Logger.LogHttpBodies).To(BeTrue())
			})
		})

		Context("NewRelic config", func() {
			It("builds default NewRelic config", func() {
				appConfig, err := config.NewAppConfig()

				Expect(err).ToNot(HaveOccurred())
				Expect(appConfig.NewRelic).ToNot(BeNil())

				Expect(appConfig.NewRelic.AppName).To(Equal(""))
				Expect(appConfig.NewRelic.LicenceKey).To(Equal(""))
				Expect(appConfig.NewRelic.Enabled).To(Equal(false))
				Expect(appConfig.NewRelic.LabelEnvironment).To(Equal("local"))
				Expect(appConfig.NewRelic.LabelAccount).To(Equal(""))
				Expect(appConfig.NewRelic.LabelRole).To(Equal(""))
			})

			It("reads NewRelic config from environment vars", func() {
				currentNewRelicAppName := os.Getenv("NEW_RELIC_APP_NAME")
				currentNewRelicLicenceKey := os.Getenv("NEW_RELIC_LICENCE_KEY")
				currentNewRelicEnabled := os.Getenv("NEW_RELIC_ENABLED")
				currentNewRelicLabelEnv := os.Getenv("NEW_RELIC_LABEL_ENV")
				currentNewRelicLabelAccount := os.Getenv("NEW_RELIC_LABEL_ACCOUNT")
				currentNewRelicLabelRole := os.Getenv("NEW_RELIC_LABEL_ROLE")

				os.Setenv("NEW_RELIC_APP_NAME", "some-app-name")
				os.Setenv("NEW_RELIC_LICENCE_KEY", "some-licence-key")
				os.Setenv("NEW_RELIC_ENABLED", "true")
				os.Setenv("NEW_RELIC_LABEL_ENV", "tst")
				os.Setenv("NEW_RELIC_LABEL_ACCOUNT", "smartshop-services-pre")
				os.Setenv("NEW_RELIC_LABEL_ROLE", "api_colleague_discount")

				defer os.Setenv("NEW_RELIC_APP_NAME", currentNewRelicAppName)
				defer os.Setenv("NEW_RELIC_LICENCE_KEY", currentNewRelicLicenceKey)
				defer os.Setenv("NEW_RELIC_ENABLED", currentNewRelicEnabled)
				defer os.Setenv("NEW_RELIC_LABEL_ENV", currentNewRelicLabelEnv)
				defer os.Setenv("NEW_RELIC_LABEL_ACCOUNT", currentNewRelicLabelAccount)
				defer os.Setenv("NEW_RELIC_LABEL_ROLE", currentNewRelicLabelRole)

				appConfig, err := config.NewAppConfig()

				Expect(err).ToNot(HaveOccurred())
				Expect(appConfig.NewRelic).ToNot(BeNil())

				Expect(appConfig.NewRelic.AppName).To(Equal("some-app-name"))
				Expect(appConfig.NewRelic.LicenceKey).To(Equal("some-licence-key"))
				Expect(appConfig.NewRelic.Enabled).To(Equal(true))
				Expect(appConfig.NewRelic.LabelEnvironment).To(Equal("tst"))
				Expect(appConfig.NewRelic.LabelAccount).To(Equal("smartshop-services-pre"))
				Expect(appConfig.NewRelic.LabelRole).To(Equal("api_colleague_discount"))
			})
		})
	})
})
