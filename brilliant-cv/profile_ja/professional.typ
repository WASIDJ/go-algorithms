// Imports
#import "@preview/brilliant-cv:4.0.1": (
  cv-entry, cv-entry-continued, cv-entry-start, cv-section,
)


#cv-section("職務経歴")

#cv-entry-start(
  society: [XYZ 社],
  logo: image("../assets/logos/xyz_corp.png"),
  location: [サンフランシスコ, CA],
)

#cv-entry-continued(
  title: [データサイエンス責任者],
  date: [2020 - 現在],
  description: list(
    [データサイエンティストとアナリストのチームを統括し、データ駆動型戦略の策定と実装、予測モデルとアルゴリズムの開発を推進],
    [経営幹部と連携してビジネス機会を特定し、成長を促進。データガバナンス、品質、セキュリティのベストプラクティスを導入],
  ),
  tags: ("Dataiku", "Snowflake", "SparkSQL"),
)

#cv-entry(
  title: [データアナリスト],
  society: [ABC 社],
  logo: image("../assets/logos/abc_company.png"),
  date: [2017 - 2020],
  location: [ニューヨーク, NY],
  description: list(
    [SQL と Python を用いた大規模データセットの分析、チームと協力したビジネスインサイトの発見],
    [Tableau によるデータ可視化とダッシュボードの作成、AWS を用いたデータパイプラインの構築と保守],
  ),
)

#cv-entry(
  title: [データ分析インターン],
  society: [PQR 社],
  logo: image("../assets/logos/pqr_corp.png"),
  date: list(
    [2017年 夏季],
    [2016年 夏季],
  ),
  location: [シカゴ, IL],
  description: list(
    [Python と Excel を用いたデータクレンジング、処理、分析の補助。チームミーティングに参加しプロジェクト計画と実行に貢献],
  ),
)
