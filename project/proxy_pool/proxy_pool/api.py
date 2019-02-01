import json
from flask import Flask, g, abort

from proxy_pool.utils import env, new_instance

__all__ = ['app']

app = Flask(__name__)


def get_conn():
    if not hasattr(g, 'db'):
        g.db = new_instance('proxy_pool.db', env('db.type'))
    return g.db


def get_by_count(table, count):
    conn = get_conn()

    # 检查是否有这张表
    allow_table = conn.table()
    if table not in allow_table:
        abort(404)

    datas = conn.random(table_name=table, count=count)

    result = list()
    for data in datas:
        result.append(transform(data))
    return result


def transform(data):
    data.pop('created_at')
    data.pop('request_time')
    data['protocol_type'] = 'https' if data.get('protocol_type') == 2 else 'http'
    return data


@app.route('/')
def index():
    return '<div>方法列表: <ul><li>/表名</li><li>/表名/获取的数量</li></ul></div>'


@app.route('/<table>')
def get_proxy(table):
    data = get_by_count(table, 1)
    return json.dumps(data)


@app.route('/<table>/<count>')
def get_counts(table, count):
    count = int(count)
    data = get_by_count(table, count)
    return json.dumps(data)


if __name__ == '__main__':
    app.run()
