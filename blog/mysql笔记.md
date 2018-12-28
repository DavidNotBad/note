## 获取mysql服务版本

```mysql
SELECT version() as version;
```

## 查看表描述

```mysql
SELECT
    COLUMN_NAME,
    COLUMN_COMMENT
FROM
    INFORMATION_SCHEMA. COLUMNS
WHERE
    TABLE_NAME = '表名'
    # AND table_schema = '数据库名';
```

## 查看表的主键

```mysql
SELECT
        COLUMN_KEY,
        COLUMN_NAME
FROM
        INFORMATION_SCHEMA. COLUMNS
WHERE
        table_name = 'tp_samples'
AND COLUMN_KEY = 'PRI';
```

