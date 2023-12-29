---
layout: page
title: ソフトウェア学習
permalink: /_post-of-software-study/
---

このページはソフトウェアについての学習を段階的に行うことができます

全体の構成は以下のようになっています。

- GitとGithubについて
- HTMLとCSS
- TypeScript
- Vue3
- E2Eテスト
- Golang
- APIテスト
- Database

それぞれのコンテンツは以下のようなフォーマットで作成されています
1. 学習の対象者
   1. 前提条件 
2. ゴール
3. ドキュメントのスコープ 
4. 何を教えようとしているのか 
5. ドキュメントを読み込むことで得られるメリット、デメリット

<h1>コンテンツ</h1>
<!-- タグでソフトウェア学習を設定したPOSTだけここのコンテンツで表示される様に設定している -->
{% for study-content in site.tags.software-study %}
  <ul>
   <li><a href="{{ study-content.url }}">{{ study-content.title }}</a></li>
  </ul>
{% endfor %}