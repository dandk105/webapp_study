# これは何？
ここのディレクトリでは、データベースの設計に関するドキュメントを管理します。

## ファイルの説明
- `init.sql`: データベースの初期化を行うためのSQL文を記述します。

### 環境
今回は、PostgreSQLを使用します。
PostgreSQLの詳細については、[公式ドキュメント](https://www.postgresql.jp/document/12/html/index.html)を参照してください。

### schemaについて
このdiagarmは、[mermaid](https://mermaid-js.github.io/mermaid/#/)記法を使用しています。


```mermaid
erDiagram
    USERS o{--o{ RESERVATIONS : reserved
    USERS {
        UUID id PK "UNIQUE NOT NULL "
        text name "min(1) , max(50)"
        birth_day date
    }
    RESERVATIONS 
    RESERVATIONS {
       UUID id  PK "UNIQUE NOT NULL"
       UUID reserved_user FK "NOT NULL"
       UUID reserved_room FK "NOT NULL"
       timestamp start_date  "NOT NULL"
       timestamp end_date "NOT NULL"
    }
    
    ROOMS o|--o| RESERVATIONS : contains
    ROOMS{  
        UUID id PK "UNIQUE NOT NULL "
        name text "NOT NULL"
        capacity smallint "NOT NULL"
    }
```