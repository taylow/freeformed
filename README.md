<p align="center">
  <a href="" rel="noopener">
 <img width=300px src="https://raw.githubusercontent.com/taylow/freeformed/main/docs/img/freeformed-logo.png.png" alt="Freeformed logo"></a>
</p>

<h3 align="center">Freeformed - An open source HTML form processing service!</h3>

<div align="center">
    [![Status](https://img.shields.io/badge/status-active-success.svg)]()
    [![GitHub Issues](https://img.shields.io/github/issues/taylow/freeformed.svg)](https://github.com/taylow/freeformed/issues)
    [![GitHub Pull Requests](https://img.shields.io/github/issues-pr/taylow/freeformed.svg)](https://github.com/taylow/freeformed/pulls)
    <!-- [![License](https://img.shields.io/badge/license-CC--BY--NC--SA--4.0-blue)](/LICENSE) -->
</div>

---

Note: This project is very early in development and should not be used in production until this message is removed.

## Features

✅ Implemented

- 💾 URL-encoded form data handling and storage - Handle and store arbitrary form data without having to pre-configure the expected data
- 🗄️ Multipart-form data and file handling and storage - Handle and store multipart-form with or without attached files
- 🗄️ JSON handling and storage - Handle and store json data
- Postgres storage - Store data to a postgres database
- S3 storage - Store files to an S3-compatible object store, such as MinIO, etc.

💡 Upcoming

- 🕹️ Dashboard - Create/edit forms, view/delete submissions, configure data sources, all from the comfort of your browser
- 🖥️ CLI/TUI - Host a minimal version of the form through CLI/TUI for quick and easy form processing
- Docker Compose stack - A full-feature docker-compose stack
- Supabase/Firebase/Pocketbase Support - For quick and easy hosting without the complexity of hosting an entire docker stack  
- 📧 Email forwarding - Forward form submissions directly to various email addresses
- 🪝 3rd party integrations - Forward your data to various sources, such as Slack, Discord, and Webhooks
- 🏁 Origin protection - Prevent other sites and bots from using your form URL
- 🤖 Spam protection - Use a mix of methods such as hidden field, origin, and ML to detect and categorise spam form submissions
- 🏗️ Form Builder/Expected Data - Build a form before wth the built in form builder and prevent data outside defined structure
- 🎨 View Builder - Choose from a range of views, from the default data grid, to customised charts

## How it Works

Simply use the generated URL in your form's `action` attribute, along with the `POST` method and the data and files will be stored in the storage of choice.

When you submit data through your form, the data will be available via a simple, yet customisable, dashboard.

## Stack

Freeformed is written in Go + HTMX, with SQL, HTML, and Tailwind to support.
