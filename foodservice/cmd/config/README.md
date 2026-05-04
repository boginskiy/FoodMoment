### Особенности

Приоритет подгрузки config в порядке от более значительного к менее значительному
ENV -> CLI -> FILE JSON -> default

Это значит, что, если какой-нибудь атрибут был загружен с config ENV, то config FILE JSON
не повлияет на этот атрибут.


Пример:
//  serverAddr := cfg.GetString("server_addr", "localhost:8080")
// 	dbAddr := cfg.GetString("db_addr", "postgres:5432")
// 	redisAddr := cfg.GetString("redis_addr", "redis:6379")
// 	newField := cfg.GetString("new_field", "default")