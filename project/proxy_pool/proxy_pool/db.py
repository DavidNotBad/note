import random
import sqlite3
import threading
from abc import ABC, abstractmethod
from sqlite3 import ProgrammingError
from proxy_pool.utils import env, get_now_time, is_number

lock = threading.Lock()


class DB(ABC):
    @abstractmethod
    def create_table(self, table_name):
        pass

    @abstractmethod
    def count(self, table_name):
        pass

    @abstractmethod
    def is_exists(self, table_name, ip, port, protocol_type):
        pass

    @abstractmethod
    def insert(self, table_name, data):
        pass

    @abstractmethod
    def update(self, table_name, ip, port, protocol_type, request_time):
        pass

    @abstractmethod
    def delete(self, table_name, ip, port, protocol_type):
        pass

    @abstractmethod
    def random(self, table_name, count=1):
        pass

    @abstractmethod
    def table(self):
        pass

    def save(self, table_name, data: dict, insert_callback=None, delete_callback=None):
        try:
            global lock
            lock.acquire(True)

            ip = data.get('ip')
            port = data.get('port')
            protocol_type = data.get('protocol_type')
            request_time = data.get('request_time')

            # 插入或修改数据
            check_rquest_time = request_time and is_number(request_time) and (
                        float(request_time) < float(env('PROXY_MAX_REQUEST_TIME')))
            if not self.is_exists(table_name=table_name, ip=ip, port=port, protocol_type=protocol_type):
                if check_rquest_time:
                    is_overflow = False
                    if insert_callback is not None:
                        if insert_callback() is False:
                            is_overflow = True
                    # 数据库不存在该ip, 添加到数据库
                    if is_overflow is False:
                        self.insert(table_name, data)
            else:
                if check_rquest_time:
                    # 数据库中已经存在, 更新request_time字段
                    self.update(table_name=table_name, ip=ip, port=port, protocol_type=protocol_type,
                                request_time=request_time)
                else:
                    if delete_callback is not None:
                        delete_callback()
                    # ip地址不可用, 从库中删除
                    self.delete(table_name=table_name, ip=ip, port=port, protocol_type=protocol_type)
        finally:
            lock.release()


class SQLite(DB):
    def __init__(self):
        self.conn = sqlite3.connect(env('db.file_path'), check_same_thread=False)
        self.db = self.conn.cursor()

    def create_table(self, table_name):
        sql = '''CREATE TABLE IF NOT EXISTS '{table_name}'(
            id INTEGER PRIMARY KEY AUTOINCREMENT    NOT NULL,
            ip             VARCHAR(15)   NOT NULL DEFAULT '',
            port           VARCHAR(8)   NOT NULL DEFAULT '',
            request_time   VARCHAR(16) NOT NULL DEFAULT '',
            protocol_type  VARCHAR(10) NOT NULL DEFAULT '',
            created_at      char(16) NOT NULL default '',
            updated_at     char(16) NOT NULL default ''
            );'''
        sql = sql.format(sql, table_name=table_name)
        self.db.execute(sql)
        self.conn.commit()

    def insert(self, table_name, data: dict):
        sql = """INSERT INTO '{table_name}' ({fields}) VALUES ({values})"""
        fields = "'" + ("','".join(data.keys())) + "'"
        values = "'" + ("','".join(data.values())) + "'"

        try:
            sql = sql.format(sql, table_name=table_name, fields=fields, values=values)
            self.db.execute(sql)
        except ProgrammingError:
            pass

    def select(self, table_name, limit=None):
        keys = ('ip', 'port', 'request_time', 'protocol_type', 'created_at')
        limit_sql = 'LIMIT {limit}'.format(limit=str(limit)) if limit else ''

        select = "`" + "`,`".join(keys) + "`"
        sql = "SELECT {select} FROM '{table_name}' ORDER BY request_time ASC {limit}"
        sql = sql.format(table_name=table_name, limit=limit_sql, select=select)
        curser = self.db.execute(sql)

        for row in curser:
            yield dict(map(lambda x, y: [x, y], keys, row))

    def update(self, table_name, ip, port, protocol_type, request_time):
        updated_at = get_now_time()
        sql = "UPDATE {table_name} set 'updated_at' = '{updated_at}', 'request_time'='{request_time}' "
        sql += "where ip='{ip}' and port='{port}' and protocol_type='{protocol_type}'"
        sql = sql.format(ip=ip, table_name=table_name, updated_at=updated_at, request_time=request_time, port=port, protocol_type=protocol_type)
        return self.db.execute(sql)

    def delete(self, table_name, ip, port, protocol_type):
        sql = "DELETE FROM {table_name} where ip='{ip}' and port='{port}' and protocol_type='{protocol_type}'"
        sql = sql.format(ip=ip, table_name=table_name, port=port, protocol_type=protocol_type)
        return self.db.execute(sql)

    def is_exists(self, table_name, ip, port, protocol_type):
        sql = "SELECT COUNT(*) FROM '{table_name}' WHERE ip = '{ip}' AND port = '{port}' AND protocol_type = '{protocol_type}'"
        sql = sql.format(table_name=table_name, ip=ip, port=port, protocol_type=protocol_type)
        curser = self.db.execute(sql)
        for row in curser:
            return bool(row[0])

    def save(self, table_name, data: dict, insert_callback=None, delete_callback=None):
        super().save(table_name=table_name, data=data, insert_callback=insert_callback, delete_callback=delete_callback)
        self.conn.commit()

    def count(self, table_name):
        sql = "SELECT COUNT(*) FROM '{table_name}'"
        sql = sql.format(table_name=table_name)
        curser = self.db.execute(sql)
        for row in curser:
            return row[0]

    def random(self, table_name, count=1):
        count = int(count)

        # 获取所有代理的id列表
        sql = "SELECT id FROM '{table_name}'"
        sql = sql.format(table_name=table_name)
        curser = self.db.execute(sql)
        datas = list()
        for row in curser:
            datas.append(str(row[0]))

        # 随机选取id, 根据id来获取数据
        ids = ','.join(random.sample(datas, count) if len(datas) > count else datas)
        keys = ('ip', 'port', 'request_time', 'protocol_type', 'created_at')
        select = "`" + "`,`".join(keys) + "`"
        sql = "SELECT {select} FROM '{table_name}' WHERE id in({ids})"
        sql = sql.format(table_name=table_name, select=select, ids=ids)
        curser = self.db.execute(sql)

        # 返回指定个数的随机ip
        result = list()
        for row in curser:
            result.append(dict(map(lambda x, y: [x, y], keys, row)))
        return result

    def table(self):
        sql = "SELECT name FROM 'sqlite_sequence'"
        curser = self.db.execute(sql)

        result = list()
        for row in curser:
            result.append(row[0])
        return result

    def __del__(self):
        self.conn.commit()
        self.conn.close()
