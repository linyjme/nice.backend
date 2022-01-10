## client项目路由

 AppController {}: +20ms
 {, GET} route +2ms
 AuthController {/auth}: +0ms
 {/auth/admin, GET} route +1ms
 {/auth/admin, PUT} route +1ms
 {/auth/login, POST} route +0ms
 {/auth/check, POST} route +1ms
 {/auth/renewal, POST} route +1ms
 OptionController {/option}: +1ms
 {/option, GET} route +1ms
 {/option, PUT} route +1ms
 AnnouncementController {/announcement}: +0ms
 {/announcement, GET} route +0ms
 {/announcement, POST} route +1ms
 {/announcement, DELETE} route +0ms
 {/announcement/:id, PUT} route +0ms
 {/announcement/:id, DELETE} route +1ms
 TagController {/tag}: +0ms
 {/tag, GET} route +0ms
 {/tag, POST} route +0ms
 {/tag, DELETE} route +1ms
 {/tag/:id, PUT} route +0ms
 {/tag/:id, DELETE} route +0ms
 SyndicationController {/syndication}: +1ms
 {/syndication/sitemap, GET} route +0ms
 {/syndication/rss, GET} route +0ms
 {/syndication, PATCH} route +1ms
 CategoryController {/category}: +0ms
 {/category, GET} route +0ms
 {/category, POST} route +1ms
 {/category, DELETE} route +0ms
 {/category/:id, GET} route +0ms
 {/category/:id, PUT} route +1ms
 {/category/:id, DELETE} route +0ms
 ArticleController {/article}: +1ms
 {/article, GET} route +0ms
 {/article, POST} route +0ms
 {/article, PATCH} route +1ms
 {/article, DELETE} route +1ms
 {/article/:id, GET} route +3ms
 {/article/:id, PUT} route +4ms
 {/article/:id, DELETE} route +2ms
 CommentController {/comment}: +0ms
 {/comment, GET} route +1ms
 {/comment, POST} route +0ms
 {/comment, PATCH} route +1ms
 {/comment, DELETE} route +0ms
 {/comment/:id, GET} route +0ms
 {/comment/:id, PUT} route +1ms
 {/comment/:id, DELETE} route +1ms
 LikeController {/like}: +0ms
 {/like/site, PATCH} route +1ms
 {/like/comment, PATCH} route +0ms
 {/like/article, PATCH} route +1ms
 ExpansionController {/expansion}: +0ms
 {/expansion/statistic, GET} route +0ms
 {/expansion/uptoken, GET} route +1ms
 {/expansion/google-token, GET} route +0ms
 {/expansion/database-backup, PATCH} route +0ms
 
 
## API Document

语法释义：

```ts
@JwtAuth

Get '/path/:id' ({ param? }): <ResultDataType>
```

- `Get` method, 请求类型
- `'/path/:id'` route param, 路由及路由参数
- `({ param? })` query params | body, 请求参数或提交数据体，两者互斥
- `<ResultDataType>` result data type, 请求返回的数据类型
- `@JwtAuth` required auth, 是否需要鉴权（默认为普通用户/无需鉴权，显式声明即为需要鉴权通过才可访问）

---

#### Root
  ```ts
  Get (): <APP_INFO>
  ```

#### Auth (auth)
  ```ts
  Get '/admin' (): <Auth>
  ```
  ```ts
  @JwtAuth

  Put '/admin' (Auth): <Auth>
  ```
  ```ts
  Post '/login' (AuthLogin): <ITokenResult>
  ```
  ```ts
  @JwtAuth

  Post '/check' (): <void>
  Post '/renewal' (): <ITokenResult>
  ```

#### Option (option)
  ```ts
  Get (): <Option>
  ```
  ```ts
  @JwtAuth

  Put (Option): <Option>
  ```

#### Announcement (announcement)
  ```ts
  Get ({ <common>?, state?, keyword? }): <Announcement[]>
  ```
  ```ts
  @JwtAuth

  Post (Announcement): <Announcement>
  ```
  ```ts
  @JwtAuth

  Delete (DelAnnouncements): <MongooseOpResult>
  ```
  ```ts
  @JwtAuth

  Put '/:id' (Announcement): <Announcement>
  ```
  ```ts
  @JwtAuth

  Delete '/:id' (): <MongooseOpResult>
  ```

#### Article (article)
  ```ts
  Get ({ <common>?, date?, public?, origin?, cache?, tag?, category?, keyword?, tag_slug?, category_slug? }): <Article[]>
  ```
  ```ts
  Get '/:id' (): <Article>
  ```
  ```ts
  @JwtAuth

  Put '/:id' (Article): <Article>
  ```
  ```ts
  @JwtAuth

  Delete '/:id' (): <MongooseOpResult>
  ```
  ```ts
  @JwtAuth

  Post (Article): <Article>
  ```
  ```ts
  @JwtAuth

  Patch (PatchArticles): <MongooseOpResult>
  ```
  ```ts
  @JwtAuth

  Delete (DelArticles): <MongooseOpResult>
  ```

#### Category (category)
  ```ts
  Get ({ <common>? }): <Category[]>
  ```
  ```ts
  Get '/:id' (): <Category[]>
  ```
  ```ts
  @JwtAuth

  Put '/:id' (Category): <Category>
  ```
  ```ts
  @JwtAuth

  Delete '/:id' (): <MongooseOpResult>
  ```
  ```ts
  @JwtAuth

  Post (Category): <Category>
  ```
  ```ts
  @JwtAuth

  Delete (DelCategories): <MongooseOpResult>
  ```

#### Tag (tag)
  ```ts
  Get ({ <common>?, cache?, keyword? }): <Tag[]>
  ```
  ```ts
  @JwtAuth

  Post (Tag): <Tag>
  ```
  ```ts
  @JwtAuth

  Put '/:id' (Tag): <Tag>
  ```
  ```ts
  @JwtAuth

  Delete '/:id' (): <MongooseOpResult>
  ```
  ```ts
  @JwtAuth

  Delete (DelTags): <MongooseOpResult>
  ```

#### Comment (comment)
  ```ts
  Get ({ <common>?, state?, post_id?, keyword? }): <Comment[]>
  ```
  ```ts
  Post (CreateCommentBase): <Comment>
  ```
  ```ts
  @JwtAuth

  Patch (PatchComments): <MongooseOpResult>
  ```
  ```ts
  @JwtAuth

  Delete (DelComments): <MongooseOpResult>
  ```
  ```ts
  @JwtAuth

  Get '/:id' (): <Comment>
  ```
  ```ts
  @JwtAuth

  Put '/:id' (Comment): <Comment>
  ```
  ```ts
  @JwtAuth

  Delete '/:id' (): <MongooseOpResult>
  ```

#### Syndication (syndication)
  ```ts
  Get '/sitemap' (): <SitemapXML>
  ```
  ```ts
  Get '/rss' (): <RSSXML>
  ```
  ```ts
  @JwtAuth

  Patch (): <void>
  ```

#### Like (like)
  ```ts
  Patch '/site' (): <boolean>
  ```
  ```ts
  Patch '/comment' (LikeComment): <boolean>
  ```
  ```ts
  Patch '/article' (LikeArticle): <boolean>
  ```

#### Expansion (expansion)
  ```ts
  Get '/statistic' (): <ITodayStatistic>
  ```
  ```ts
  @JwtAuth

  Get '/uptoken' (): <IUpToken>
  ```
  ```ts
  @JwtAuth

  Get '/google-token' (): <Credentials>
  ```
  ```ts
  @JwtAuth

  Patch '/database-backup' (): <void>
  ```

