#### nice.tb_administrator 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint unsigned | PRI | NO | auto_increment |  |
| 2 | created_at |  | datetime |  | YES |  |  |
| 3 | updated_at |  | datetime |  | YES |  |  |
| 4 | deleted_at |  | datetime | MUL | YES |  |  |
| 5 | uid |  | varchar(128) |  | YES |  |  |
| 6 | account |  | varchar(128) |  | NO |  |  |
| 7 | password |  | varchar(256) |  | YES |  |  |
| 8 | nick_name |  | varchar(128) |  | YES |  |  |
| 9 | header_img | 用户头像 | varchar(128) |  | YES |  | http://qmplusimg.henrongyi.top/head.png |
| 10 | email |  | varchar(256) |  | YES |  |  |
| 11 | phone |  | varchar(32) |  | YES |  |  |
| 12 | company |  | varchar(256) |  | YES |  |  |
| 13 | address |  | varchar(256) |  | YES |  |  |
| 14 | remark |  | varchar(512) |  | YES |  |  |
| 15 | status |  | tinyint unsigned |  | YES |  | 0 |
