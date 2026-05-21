// Imports
#import "@preview/brilliant-cv:4.0.1": cv-entry, cv-section, h-bar


#cv-section("学歴")

#cv-entry(
  title: [データサイエンス修士],
  society: [カリフォルニア大学ロサンゼルス校],
  date: [2018 - 2020],
  location: [アメリカ],
  logo: image("../assets/logos/ucla.png"),
  description: list(
    [論文: 機械学習アルゴリズムとネットワーク分析による通信業界の顧客離反予測],
    [課程: ビッグデータシステムと技術 #h-bar() データマイニングと探索 #h-bar() 自然言語処理],
  ),
)

#cv-entry(
  title: [コンピューターサイエンス学士],
  society: [カリフォルニア大学ロサンゼルス校],
  date: [2014 - 2018],
  location: [アメリカ],
  logo: image("../assets/logos/ucla.png"),
  description: list(
    [論文: 機械学習による株価予測の検討: 回帰モデルと時系列モデルの比較研究],
    [課程: データベースシステム #h-bar() コンピューターネットワーク #h-bar() ソフトウェア工学 #h-bar() 人工知能],
  ),
)
