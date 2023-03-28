**SOURCE CODE WILL BE RELEASED AT 10 STARS!**

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

  <h3 align="center">Klimt Stealer</h3>

  <p align="center">
The most powerful credential & information stealing tool written in <a href="https://go.dev">GoLang</a>
    <br />
Created and maintained by <a href="https://discord.com/users/900072916597735444">yahweh#7023</a>
    <br />
    <br />
    <a href="https://github.com/codeuk/klimt">View Demo</a>
    ·
    <a href="https://github.com/codeuk/klimt/issues">Report Bug</a>
    ·
    <a href="https://github.com/codeuk/klimt/issues">Request Feature</a>
  </p>
</div>
<br />

# About This Project

Klimt is an offensive security tool designed for credential stealing, system information analysis and security assessments. It features a range of options for compromising different types of systems and applications, including Discord, Various CryptoCurrency Wallets, Web Browsers, Roblox, and [many more](https://github.com/codeuk/klimt#features). Klimt also includes a persistent reverse shell option that allows users to execute commands from a remote server, along with an interactive GUI builder.


## Features
*Features with the 💎 emoji before them are held for the private/professional version only*

<details>
<summary><strong>Credential Stealing</strong></summary>
<br>

- Discord Stealer: Steals authentication tokens and account information from 30+ Discord locations.
- Wallet Stealer: Steals cryptocurrency wallet files and credentials from popular wallets.
- 💎 Browser Stealer: Steals browser cookies, passwords, history and downloads from 20+ web browsers.
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

</details>

<details>
<summary><strong>GUI Builder</strong></summary>
<br>

![image](https://user-images.githubusercontent.com/75194878/228186187-a354a6a2-e9d2-4a93-ac29-b967d4db443b.png)

![image](https://user-images.githubusercontent.com/75194878/228186569-fe86e8b7-80fa-4a03-a0a5-5df7d0206a19.png)

</details>

<details>
<summary><strong>Misc</strong></summary>
<br>

- Reverse Shell: Optional connection to your external listener server (netcat, msf, etc.)
- Custom GUI Builder with interactive widgets (Embed color picker, tabs, etc.)
- Encrypts strings stored in the config so that they aren't searchable in the binary, such as:
  - Discord Webhook URL
  - Reverse Shell Server Information

</details>

## Embed Showcase
**- Embed Images look blurred as of now, this will be fixed later!**

*Embed layouts may differ depending on the stealer configuration*

<details>
<summary><strong>Overview</strong></summary>
<br>

- Logs.zip file, Scrape numbers, Browser Credentials & Crypto Wallets overview

![image](https://user-images.githubusercontent.com/75194878/228089502-bf0be37d-2260-488a-9fdf-bf65275e75ee.png)
</details>

<details>
<summary><strong>System</strong></summary>
<br>

- Operating System, Networking & File structure information

![image](https://user-images.githubusercontent.com/75194878/228090113-9ac06188-b6cc-4d8f-b55a-3abc88213533.png)
</details>

<details>
<summary><strong>Discord</strong></summary>
<br>

- Discord paths (where tokens were found), infection status, account information derived from each token.

![image](https://user-images.githubusercontent.com/75194878/228090586-41b1506c-34b8-4514-b2f6-87c73e312b03.png)
</details>


## Installation

To install Klimt, make sure you have [GoLang](https://go.dev) installed and follow these steps:

1. Clone the repository using `git clone https://github.com/codeuk/klimt.git`
2. Open a terminal and navigate to the directory you installed Klimt in.
3. Install the required dependencies using `go get ./...`
4. Build the builder using `go build builder.go` or by running the `build.bat` file if you're on Windows

## Usage

To use the Klimt Builder, follow these steps:

1. In the Klimt directory, run `builder` or open the new `builder.exe` file to open the Klimt Stealer Builder.
2. Wait a few seconds, and when the builder GUI opens, configure it how you wish and press `Compile Stealer`.
3. If the build is successful, the executable should be in `build/agent.exe`, and can now use this stub as you wish.

## Disclaimer

Klimt is intended for legal and ethical use only. The developers and contributors of Klimt are not responsible for any illegal or unethical activities performed using this tool. Users of Klimt are solely responsible for their actions and are advised to use the tool for legitimate security testing purposes only.

## License

Klimt is licensed under the GNU General Public License v3.0. See the LICENSE file for more information.

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
