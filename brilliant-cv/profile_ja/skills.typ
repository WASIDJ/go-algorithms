// Import
#import "@preview/brilliant-cv:4.0.1": cv-section, cv-skill, h-bar


#cv-section("スキルと興味")

#cv-skill(
  type: [言語],
  info: [英語 #h-bar() フランス語 #h-bar() 中国語],
)

#cv-skill(
  type: [技術スタック],
  info: [Tableau #h-bar() Python (Pandas/Numpy) #h-bar() PostgreSQL],
)

#cv-skill(
  type: [趣味],
  info: [水泳 #h-bar() 料理 #h-bar() 読書],
)
