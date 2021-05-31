# Sync Google Calendar


**NOTE:** This README and CLI messages are written in Japanese.

このCLIアプリケーションを使うと、Googleカレンダーのスケジュール（イベント）を他のGoogleカレンダーへコピーすることが出来ます。  
個人で使っているGoogleカレンダーのスケジュールが会社のGoogleカレンダーに登録されていないために、大事な予定の時間に会議を入れられてしまった！  
そのようなことを防ぐのに役立ちます。

ただし、[client_secret.json の作成](https://developers.google.com/workspace/guides/create-credentials)が必要です。  
作成方法について説明すると長くなってしまうので、ここでは省略します。


## Installation

Go 1.16+:

```sh
go install github.com/nekonenene/sync-google-calendar@latest
```

Otherwise:

```sh
go get github.com/nekonenene/sync-google-calendar@latest
```


## Usage

### Example

```sh
sync-google-calendar --credential-file client_secret.json --start-date 2021/01/01 --end-date 2021/01/14 --title-prefix "[Private] "
```

### Parameters

You can see all parameters:

```sh
sync-google-calendar --help
```


## Build

```sh
make build
```


## License

[MIT](https://choosealicense.com/licenses/mit/)
