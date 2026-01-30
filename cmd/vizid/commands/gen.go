package commands

import (
	"fmt"

	"github.com/ryanl/vizid/internal/generator"
	"github.com/ryanl/vizid/internal/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	compYear   bool
	compMonth  bool
	compDay    bool
	compHour   bool
	compMinute bool
	compSecond bool
	compMs     bool
	compUUID   bool
	userDefined bool
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new VIZID (visual form, suitable for filenames)",
	RunE: func(cmd *cobra.Command, args []string) error {
		components := model.Components{
			Year:   viper.GetBool("components.year"),
			Month:  viper.GetBool("components.month"),
			Day:    viper.GetBool("components.day"),
			Hour:   viper.GetBool("components.hour"),
			Minute: viper.GetBool("components.minute"),
			Second: viper.GetBool("components.second"),
			Ms:     viper.GetBool("components.ms"),
			UUID:   viper.GetBool("components.uuid"),
		}

		if userDefined {
			components = model.Components{}
			if cmd.Flags().Changed("year") {
				components.Year = compYear
			}
			if cmd.Flags().Changed("month") {
				components.Month = compMonth
			}
			if cmd.Flags().Changed("day") {
				components.Day = compDay
			}
			if cmd.Flags().Changed("hour") {
				components.Hour = compHour
			}
			if cmd.Flags().Changed("minute") {
				components.Minute = compMinute
			}
			if cmd.Flags().Changed("second") {
				components.Second = compSecond
			}
			if cmd.Flags().Changed("ms") {
				components.Ms = compMs
			}
			if cmd.Flags().Changed("uuid") {
				components.UUID = compUUID
			}
		} else {
			if cmd.Flags().Changed("year") {
				components.Year = compYear
			}
			if cmd.Flags().Changed("month") {
				components.Month = compMonth
			}
			if cmd.Flags().Changed("day") {
				components.Day = compDay
			}
			if cmd.Flags().Changed("hour") {
				components.Hour = compHour
			}
			if cmd.Flags().Changed("minute") {
				components.Minute = compMinute
			}
			if cmd.Flags().Changed("second") {
				components.Second = compSecond
			}
			if cmd.Flags().Changed("ms") {
				components.Ms = compMs
			}
			if cmd.Flags().Changed("uuid") {
				components.UUID = compUUID
			}
		}

		opts := model.Options{
			Timezone: viper.GetString("timezone"),
			Warn:     viper.GetBool("warn"),
			Custom:   viper.GetBool("custom") || userDefined,
			Components: components,
		}

		id, _, warnMsg, err := generator.Generate(opts)
		if err != nil {
			return err
		}
		if opts.Warn && warnMsg != "" {
			fmt.Println("WARN:", warnMsg)
		}
		fmt.Println(id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().BoolVarP(&userDefined, "user-defined", "u", false, "start with all components disabled; explicitly enable each component")
	genCmd.Flags().BoolVar(&compYear, "year", true, "include year")
	genCmd.Flags().BoolVar(&compMonth, "month", true, "include month")
	genCmd.Flags().BoolVar(&compDay, "day", true, "include day")
	genCmd.Flags().BoolVar(&compHour, "hour", true, "include hour")
	genCmd.Flags().BoolVar(&compMinute, "minute", true, "include minute")
	genCmd.Flags().BoolVar(&compSecond, "second", true, "include second")
	genCmd.Flags().BoolVar(&compMs, "ms", true, "include milliseconds")
	genCmd.Flags().BoolVar(&compUUID, "uuid", true, "include uuid")

	_ = viper.BindPFlag("custom", genCmd.Flags().Lookup("user-defined"))
}
