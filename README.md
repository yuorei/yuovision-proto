# YuoVison Protobuf

[https://yuovision.yuorei.com](https://yuovision.yuorei.com)

YuoVisionは、動画の視聴とアップロードが可能な動画配信プラットフォームです。
あなたの動画をアップロードして、世界中の人々と共有しましょう。新しい発見をYuoVisionでシェアし、一緒に新しい世界を探求しましょう。

## 概要
本リポジトリはYuoVisonに関わるprotobufを管理しています。
現在マイクロサービスに移行中です。

## 運用方法
本リポジトリをsubmoduleとして追加することで運用しています。

### 工夫したこと

GitHub Actionsでmainにpushされた時にprptobufからGoのコードを生成します。
