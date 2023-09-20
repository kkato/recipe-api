# recipe-api
HTTPエンドポイントを作成する。
- 以下のエンドポイントを実装する。
    - `POST /recipes` -> レシピを作成
    - `GET /recipes` -> 全レシピ一覧を返す
    - `GET /recipes/{id}` -> 指定レシピ一つを返す
    - `PATCH /recipes/{id}` -> 指定レシピを更新
    - `DELETE /recipes/{id}` -> 指定レシピの削除
- レスポンスは全てJSON形式で返す。
- 上記エンドポイントに対するHTTPレスポンスステータスコードは、すべて`200`とする。
- 上記以外のエンドポイントに対するHTTPレスポンスステータスコードは、すべて`404`とする。

#### `POST /recipes`エンドポイント
- `recipe`を新規作成する。
- 期待する`request`形式: `POST /recipes`
    - `title`, `making_time`, `servers`, `ingredients`, `cost`
    - 上記パラメータはすべて必須。
- 期待する`response`形式:

成功`response`
```json
{
    "message": "Recipe successfully created!",
    "recipe": [
        {
            "id": 3,
            "title": "トマトスープ",
            "making_time": "15分",
            "serves": "5人",
            "ingredients": "玉ねぎ,トマト,スパイス、水",
            "cost": "450",
            "created_at": "2016-01-12 14:10:12",
            "updated_at": "2016-01-12 14:10:12"
        }
    ]
}
```

失敗`response`
```json
{
    "message": "Recipe creation failed",
    "required": "title, making_time, serves, ingredients, cost"
}
```

#### `GET /recipes`エンドポイント
- データベースのすべてのレシピを返す。
- 期待する`request`形式: `GET /recipes`
- 期待する`response`形式:
```json
{
    "recipes": [
        {
            "id": 1,
            "title": "チキンカレー",
            "making_time": "45分",
            "servers": "4人",
            "ingredients": "玉ねぎ,肉,スパイス",
            "cost": "1000"
        },
        {
            "id": 2,
            "title": "オムライス",
            "making_time": "30分",
            "servers": "2人",
            "ingredients": "玉ねぎ,卵,スパイス,醤油",
            "cost": "1000" 
        },
        {
            "id": 3,
            "title": "トマトスープ",
            "making_time": "15分",
            "servers": "5人",
            "ingredients": "玉ねぎ,トマト,スパイス,水",
            "cost": "450"
        }
    ]
}
```

#### `GET/recipes/{id}`エンドポイント
- 指定`id`のレシピのみを返します。
- 期待する`request`形式: `GET /recipes/1`
- 期待する`response`形式:
```json
{
    "message": "Recipe detrails by id"
    "recipes": [
        {
            "id": 1,
            "title": "チキンカレー",
            "making_time": "45分",
            "servers": "4人",
            "ingredients": "玉ねぎ,肉,スパイス",
            "cost": "1000"
        }
    ]
}
```

#### `PATCH /recipes/{id}`エンドポイント
- 指定`id`のレシピを更新し、更新したレシピを返します。
- 期待する`request`形式: `PATCH /recipes/{id}`
    - `Body`フィールド: `title`, `making_time`, `servers`, `ingredients`, `cost`
- 期待する`response`形式:
```json
{
    "message": "Recipe successfully updated!",
    "recipe": [
        {
            "title": "トマトスープレシピ",
            "making_time": "15分",
            "serves": "5人",
            "ingredients": "玉ねぎ,トマト,スパイス、水",
            "cost": "450"
        }
    ]
}
```

#### `DELETE /recipes/{id}`エンドポイント
- 指定`id`のレシピを削除します。
- 期待する`request`形式: `DELETE /recipes/1`
- 期待する`response`形式:

成功:
```json
{   "message": "Recipe successfully removed!"   }
```

失敗(指定`id`のレシピが存在しない場合):
```json
{   "message": "No Recipe found"   }
``` 