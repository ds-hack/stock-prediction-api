# stock-prediction-api

当プロジェクトは、様々なデータ分析案件で利用できる分析システムの枠組みを作り上げることを最終目標としています。

株価予測ダッシュボードの構築という具体的なサンプルをもとに、分析用データベースの構築・データハンドリング・モデリング・API開発・ダッシュボード開発を行います。

## プロジェクトリポジトリ

- data-infrastructure(Python) : Web APIやクローリングによる分析用データベースの構築
- stochastic-modeling(Python) : データハンドリング、特徴量作成・モデリングを行い、株価予測モデルを構築
- **stock-prediction-api(Go)** : 株価予測ダッシュボードで利用するRestful APIのバックエンド実装
- stock-prediction-dashboard(TypeScript) : 株価予測ダッシュボードのフロントエンド実装

![architecture-api](https://user-images.githubusercontent.com/56133802/76997326-7ceae080-6996-11ea-998e-2188d48707ef.png)

## Now Developing

RESTful APIの開発はスキーマ駆動開発で進めます。

- [OpenAPI 3.0 Specification](https://swagger.io/docs/specification)
- [openapi-generator](https://openapi-generator.tech/)
