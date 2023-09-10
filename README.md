# WebAppを勉強するためのレポジトリ

## フォルダ・ファイル構成

### backend/

go　言語で書かれているAPIサーバーです

Docker image にビルドします

詳細はREADME.mdまたは[wiki](https://github.com/dandk105/webapp_study/wiki/%E3%81%93%E3%81%AE%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6#frontend)を確認してください

### frontend/

vueのテンプレートアプリを利用したフロントエンドサーバーです

Docker imageにビルドします

詳細はREADME.mdまたは[wiki](https://github.com/dandk105/webapp_study/wiki/%E3%81%93%E3%81%AE%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6#frontend)を確認してください

### docer-compose.yml

backendのDocker image
frontendのDocker image

２つのDocker　imageをローカルで連動して動かすためのものです

### 開発環境について

開発環境の詳細についてはwikiを確認してください

[wiki](https://github.com/dandk105/webapp_study/wiki)