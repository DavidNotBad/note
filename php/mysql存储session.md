## 创建表

```mysql
CREATE TABLE `tc_session` (
`skey` char(32) CHARACTER SET ascii NOT NULL,
`data` text,
`expire` int(11) NOT NULL,
PRIMARY KEY (`skey`),
KEY `index_session_expire` (`expire`) USING BTREE
) ENGINE=innoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


Schema::create('session_cache', function (Blueprint $table) {
    $table->string('skey', 255)->primary();
    $table->text('data');
    $table->char('expire', 12)->index();

    $table->engine = 'innoDB';
    $table->charset = 'utf8';
    $table->collation = 'utf8_general_ci';
});
```

## 文件

```php
<?php
defined('ROOT_FW_PATH') or define('ROOT_FW_PATH', realpath(__DIR__ . '/../../'));
require_once(ROOT_FW_PATH . 'application/function/includes/DB.php');

class MySessionHandler implements SessionHandlerInterface
{
    private $db;
    private $table;
    private $ecs;
    private $fullTable;
    private $expireTime;
    private $time;

    public function __construct(cls_mysql $db, ECS $ecs, $table)
    {
        $this->db = $db;
        $this->ecs = $ecs;
        $this->table = $table;
        $this->fullTable = $this->ecs->table($table);
        $this->expireTime = get_cfg_var('session.gc_maxlifetime');
        $this->time = time();
    }

    /**
     * Close the session
     * @link http://php.net/manual/en/sessionhandlerinterface.close.php
     * @return bool <p>
     * The return value (usually TRUE on success, FALSE on failure).
     * Note this value is returned internally to PHP for processing.
     * </p>
     * @since 5.4.0
     */
    public function close()
    {
        return true;
    }


    /**
     * Cleanup old sessions
     * @link http://php.net/manual/en/sessionhandlerinterface.gc.php
     * @param int $maxlifetime <p>
     * Sessions that have not updated for
     * the last maxlifetime seconds will be removed.
     * </p>
     * @return bool <p>
     * The return value (usually TRUE on success, FALSE on failure).
     * Note this value is returned internally to PHP for processing.
     * </p>
     * @since 5.4.0
     */
    public function gc($maxlifetime)
    {
        $where = "expire <= {$maxlifetime}";
        return !! $this->db->delete($where, $this->table);
    }

    /**
     * Initialize session
     * @link http://php.net/manual/en/sessionhandlerinterface.open.php
     * @param string $save_path The path where to store/retrieve the session.
     * @param string $name The session name.
     * @return bool <p>
     * The return value (usually TRUE on success, FALSE on failure).
     * Note this value is returned internally to PHP for processing.
     * </p>
     * @since 5.4.0
     */
    public function open($save_path, $name)
    {
        return true;
    }


    /**
     * Destroy a session
     * @link http://php.net/manual/en/sessionhandlerinterface.destroy.php
     * @param string $session_id The session ID being destroyed.
     * @return bool <p>
     * The return value (usually TRUE on success, FALSE on failure).
     * Note this value is returned internally to PHP for processing.
     * </p>
     * @since 5.4.0
     */
    public function destroy($session_id)
    {
        $where = [
            'skey' => $session_id,
        ];
        return !! $this->db->delete($where, $this->table);
    }

    /**
     * Read session data
     * @link http://php.net/manual/en/sessionhandlerinterface.read.php
     * @param string $session_id The session id to read data for.
     * @return string <p>
     * Returns an encoded string of the read data.
     * If nothing was read, it must return an empty string.
     * Note this value is returned internally to PHP for processing.
     * </p>
     * @since 5.4.0
     */
    public function read($session_id)
    {
        $sql = "select `data` from {$this->fullTable} where skey='{$session_id}' and expire > {$this->time}";
        return $this->db->getOne($sql) ? : '';
    }

    /**
     * Write session data
     * @link http://php.net/manual/en/sessionhandlerinterface.write.php
     * @param string $session_id The session id.
     * @param string $session_data <p>
     * The encoded session data. This data is the
     * result of the PHP internally encoding
     * the $_SESSION superglobal to a serialized
     * string and passing it as this parameter.
     * Please note sessions use an alternative serialization method.
     * </p>
     * @return bool <p>
     * The return value (usually TRUE on success, FALSE on failure).
     * Note this value is returned internally to PHP for processing.
     * </p>
     * @since 5.4.0
     */
    public function write($session_id, $session_data)
    {
        //已存在, 修改数据
        $hasSession = $this->hasSession($session_id);
        if($hasSession) {
            $updateData = [
                'data' => $session_data
            ];
            $where = [
                'skey' => $session_id
            ];
            return !! $this->db->edit($updateData, $where, $this->table);
        }

        //删除过期的数据
        $this->deleteOutOfTime($session_id);

        //添加数据
        $insertData = [
            'skey' => $session_id,
            'data' => $session_data,
            'expire' => $this->expireTime,
        ];
        return !! $this->db->add($insertData, $this->table);
    }

    /**
     * @param mixed $table
     */
    public function setTable($table)
    {
        $this->table = $table;
    }

    private function hasSession($session_id)
    {
        return !! $this->read($session_id);
    }

    /**
     * @param mixed $expireTime
     */
    public function setExpireTime($expireTime)
    {
        $this->expireTime = $expireTime;
    }

    private function isOutOfTime($session_id)
    {
        $sql = "select count(*) from {$this->fullTable} where skey='{$session_id}' and expire <= {$this->time}";
        return !! $this->db->getOne($sql);
    }

    private function deleteOutOfTime($session_id)
    {
        //已过期, 删除数据
        $isOutOfTime = $this->isOutOfTime($session_id);
        if($isOutOfTime) {
            $where = [
                'skey' => $session_id
            ];
            return $this->db->delete($where, $this->table);
        }
        return true;
    }
}
```

## 使用

```php
ini_set("session.save_handler","user");

$mysql = DB::instance();
$ecs = DB::ecs();
$table = 'session_cache';


$handler = new MySessionHandler($mysql, $ecs, $table);
$handler->setExpireTime(time() + 999999999999);

session_set_save_handler($handler, true);
```













