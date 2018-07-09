## 权限管理 & 无限极分类

```php
/*
CREATE TABLE IF NOT EXISTS `think_access` (
  `role_id` smallint(6) unsigned NOT NULL,
  `node_id` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) NOT NULL,
  `module` varchar(50) DEFAULT NULL,
  KEY `groupId` (`role_id`),
  KEY `nodeId` (`node_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `think_node` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `sort` smallint(6) unsigned DEFAULT NULL,
  `pid` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `level` (`level`),
  KEY `pid` (`pid`),
  KEY `status` (`status`),
  KEY `name` (`name`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `think_role` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `pid` smallint(6) DEFAULT NULL,
  `status` tinyint(1) unsigned DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `status` (`status`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ;
 
CREATE TABLE IF NOT EXISTS `think_role_user` (
  `role_id` mediumint(9) unsigned DEFAULT NULL,
  `user_id` char(32) DEFAULT NULL,
  KEY `group_id` (`role_id`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
 */
class auth
{
    /**
     * @var $db cls_mysql
     */
    protected $db;
    /**
     * @var $ecs ECS
     */
    protected $ecs;
    protected $table;

    public function __construct()
    {
        $this->db = $GLOBALS['db'];
        $this->ecs = $GLOBALS['ecs'];
        $this->table = (object) array(
            'access' => $this->ecs->table('access'),
            'node' => $this->ecs->table('node'),
            'role' => $this->ecs->table('role'),
            'roleUser' => $this->ecs->table('role_user'),
            'user' => $this->ecs->table('user'),
        );

        //尝试开启session
        isset($_SESSION) || session_start();
    }


    /**
     * 获取用户权限列表
     * @param int $userId   用户id
     * @param array $select    获取的字段
     * @param array $where      条件
     * @param string $order     排序
     * @param string $limit     分段
     * @return array|bool
     */
    public function access($userId, $select=array('node.*'), array $where=array(), $order='level asc, sort desc', $limit=null)
    {
        $table = implode(',', array(
            "{$this->table->roleUser} as role_user",
            "{$this->table->access} as access",
            "{$this->table->node} as node",
        ));

        $whereCondition = array(
            "role_user.user_id = {$userId}",
            'role_user.role_id = access.role_id',
            'access.node_id = node.id',
            'node.`status` = 1',
        );
        $where = count($where) ? array_merge($whereCondition, $where) : $whereCondition;
        $where = implode(' and ', $where);

        $select = implode(',', $select);
        $order = is_null($order) ? '' : 'order by ' . $order;
        $limit = is_null($limit) ? '' : 'limit ' . $limit;

        $sql = "select {$select} from {$table} where {$where} {$order} {$limit}";
        return $this->db->getAll($sql);
    }

    /**
     * 获取下级权限列表
     * @param int $userId       用户id
     * @param int $level        父级id
     * @param array $select     获取的字段
     * @param string $order     排序
     * @param string $limit     分段
     * @return array|bool
     */
    public function childrenAccess($userId, $level, $select=array('node.*'), $order='level asc, sort desc', $limit=null)
    {
        return $this->access($userId, $select, array(
            "node.pid = {$level}"
        ), $order, $limit);
    }

    /**
     * 获取多级权限列表
     * @param int $userId       用户id
     * @param int $level        最大级别
     * @param array $select     获取的字段
     * @param string $order     排序
     * @param string $limit     分段
     * @return array|bool
     */
    public function childrensAccess($userId, $level, $select=array('node.*'), $order='level asc, sort desc', $limit=null)
    {
        return $this->access($userId, $select, array(
            "node.level <= {$level}"
        ), $order, $limit);
    }

    /**
     * 无限极分类关系树
     * @param array $items        源数组
     * @param string $sign  附加键名
     * @param string $id    id名
     * @param string $pid   pid名
     * @return array        关系树
     */
    protected function tree(array $items, $sign='_children', $id='id', $pid='pid')
    {
        $items = array_combine(array_column($items, $id), $items);
        foreach ($items as $item)
        {
            $items[$item[$pid]][$sign][$item[$id]] = &$items[$item[$id]];
        }

        return isset($items[0][$sign]) ? $items[0][$sign] : array();
    }

    /**
     * 获取二维数组特定的列
     * @param array $array  原数组
     * @param array $keys   获取的列名
     * @return array        筛选后的二维数组
     */
    protected function array_columns(array $array, array $keys)
    {
        return array_map(function($item)use($keys){
            extract($item);
            return compact($keys);
        }, $array);
    }


}		
```