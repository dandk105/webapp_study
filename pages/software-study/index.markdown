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
3. ゴール
3. ドキュメントのスコープ 
4. 何を教えようとしているのか 
5. ドキュメントを読み込むことで得られるメリット、デメリット

<!--ここではソフトウェア学習のコンテンツとして投稿した内容を表示する様にしたい-->
{% for tag in site.tags %}
<h3>{{ tag[0] }}</h3>
  <ul>
    {% for post in tag[1] %}
      <li><a href="{{ post.url }}">{{ post.title }}</a></li>
    {% endfor %}
  </ul>
{% endfor %}