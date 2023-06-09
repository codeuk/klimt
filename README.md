<div align="center">
  <a href="https://github.com/codeuk/klimt">
    <img src="https://cdn4.iconfinder.com/data/icons/internet-security-flat-2/32/Internet_Security_danger_malware_website_web_virus-512.png" alt="Logo" width="130" height="130">
  </a>

<h1 align="center">Klimt Stealer</h3>
  <p align="center">
  <a href="https://github.com/codeuk/klimt/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/codeuk/klimt.svg?style=flat-square">
  </a>
  <a href="https://github.com/codeuk/klimt/releases">
    <img src="https://img.shields.io/badge/release-v0.1 alpha-blue.svg?style=flat-square">
  </a>
  <a href="https://github.com/codeuk/klimt/blob/master/LICENSE.txt">
    <img src="https://img.shields.io/badge/license-GNU%20v3-yellow.svg?style=flat-square">
  </a>
  <a href="https://github.com/codeuk/klimt/stargazers">
    <img src="https://img.shields.io/github/stars/codeuk/klimt">
  </a>
  <a href="https://opensource.org">
    <img src="https://img.shields.io/badge/open%20source-%E2%9D%A4-brightgreen.svg?style=flat-square">
  </a>
  </p>

  <p align="center">
The most powerful credential & information stealing tool, written in <a href="https://go.dev">GoLang</a>
    <br />
Created and maintained by <a href="https://discord.com/users/1051286621414244372">codeuk</a>
    <br />
    <br />
    <a href="https://github.com/codeuk/klimt#features">View Features</a>
    ·
    <a href="https://github.com/codeuk/klimt#embed-showcase">Embed Showcase</a>
    ·
    <a href="https://github.com/codeuk/klimt#installation">Installing/Usage</a>
  </p>
</div>
<br />

# About This Project
Klimt is an offensive security tool designed for credential stealing, system information analysis and security assessments. It features a range of options for compromising different types of systems and applications, including Discord, Various CryptoCurrency Wallets, Web Browsers, Roblox, and [many more](https://github.com/codeuk/klimt#features). Klimt also includes a persistent reverse shell option that allows users to execute commands from a remote server, along with an interactive GUI builder.

Note: **When using this program, do not upload the stub (build/agent.exe) to any virus scanning platforms (ex. VirusTotal) as this will only increase detections in the future.**

## Features
*Features with the 💎 emoji before them are held for the private/professional version only*

<details>
<summary><strong>Credential Stealing</strong></summary>
<br>

- Discord Stealer: Steals authentication tokens and account information from 30+ Discord locations.
- Wallet Stealer: Steals cryptocurrency wallet files and credentials from popular wallets.
- FileZilla Stealer: Steals recent server credentials and configuration settings from the FileZilla client.
- Browser Stealer: Steals browser cookies, passwords, history, downloads and credit cards from Web Browsers.
- 💎 Steam Stealer: Steals account authentication tokens from the Steam client.
- 💎 Telegram Stealer: Steals client session data from the Telegram client.
- 💎 Roblox Stealer: Steals login credentials and authentication tokens from the Roblox game client & website.

</details>

<details>
<summary><strong>Program Injection</strong></summary>
<br>

- Discord: Injects a custom Javascript package into the Discord process for added functionality.
- 💎 Startup: Adds Klimt to the list of startup programs to ensure persistence.
- 💎 Browsers: Injects a custom module into major web browsers for added functionality.

</details>

<details>
<summary><strong>System Information</strong></summary>
<br>

- General: Collects general system information using the Windows registry and WMIC.
- Local Files: Scans the target machines files.
- Installed Software: Collects information on installed software.
- Network Connections: Logs network activity and connections.
- 💎 WiFi/SSID Credentials: Steals login credentials and SSID information from all available WiFi connections.

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
*Embed layouts may differ depending on the stealer configuration and version (below is the free version).*

<img src="https://user-images.githubusercontent.com/75194878/230427142-90f3887b-f1ab-4bbe-bd21-877b42432cc9.png" alt="Overview" width=47% height=47%><img src="https://user-images.githubusercontent.com/75194878/228549045-6c3a4fac-20fc-40b8-811c-3d4722bccf49.png" alt="System" width=49% height=49%><img src="https://user-images.githubusercontent.com/75194878/228552046-478d6acb-fe4f-4f4b-9f50-7e5bbdea93a5.png" alt="System" width=50% height=50%>

## Builder Showcase
*Layout and theme of the builder may change in the future*


<img src="https://user-images.githubusercontent.com/75194878/228554304-e88ba688-a1d2-4042-bd78-e8bf8b1f224f.png" alt="Overview" width=50% height=50%><img src="https://user-images.githubusercontent.com/75194878/228554522-0a717830-1ff3-4112-b1e1-e56a5600774c.png" alt="Overview" width=50% height=50%>

## Installation

To install Klimt, make sure you have [GoLang](https://go.dev) and [GCC](https://sourceforge.net/projects/mingw-w64/) installed and follow these steps:

***If you can't install GCC, just use the pre-compiled builder in the [releases](https://github.com/codeuk/klimt/releases) page instead of doing the following.***


1. Clone the repository using `git clone https://github.com/codeuk/klimt.git`
2. Open a terminal and navigate to the directory you installed Klimt in.
3. Download all pre-requisites by running `go get ./...`
4. Build the builder using `go build builder.go` or by running `build.bat`

If you encounter any errors during this process, first look it up and see if it's an easy fix, and if you still can't manage to fix it, create an <a href="https://github.com/codeuk/klimt/issues">issue</a>

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
