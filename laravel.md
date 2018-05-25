# laravel

打印sql日志

```php
//将以下代码复制到AppServiceProvider中的boot方法中
protected function toSqlLog()
{
    $filePath = storage_path('logs'.DIRECTORY_SEPARATOR.'sql' . DIRECTORY_SEPARATOR);
    is_dir($filePath) || @mkdir($filePath, 0777, true);
    $logFile = $filePath . date('Y-m-d') . '.log';

    DB::listen(
        function ($sql)use($logFile) {
            // $sql is an object with the properties:
            //  sql: The query
            //  bindings: the sql query variables
            //  time: The execution time for the query
            //  connectionName: The name of the connection

            // To save the executed queries to file:
            // Process the sql and the bindings:
            foreach ($sql->bindings as $i => $binding) {
                if ($binding instanceof \DateTime) {
                    $sql->bindings[$i] = $binding->format('\'Y-m-d H:i:s\'');
                } else {
                    if (is_string($binding)) {
                        $sql->bindings[$i] = "'$binding'";
                    }
                }
            }

            // Insert bindings into query
            $query = str_replace(array('%', '?'), array('%%', '%s'), $sql->sql);

            $query = vsprintf($query, $sql->bindings);

            // Save the query to file
            $data = '['.date('Y-m-d H:i:s') . '] ' . $query;
            file_put_contents($logFile, $data.PHP_EOL, FILE_APPEND);
        }
    );
}
```



