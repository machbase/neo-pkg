'use strict';

const {Client} = require('machcli');

const dbConf = {
    host: '127.0.0.1',
    port: 5656,
    user: 'sys',
    password: 'manager',
};

const db = new Client(dbConf);
const conn = db.connect();
const rows = conn.query("SELECT NAME, VALUE FROM V$SYSSTAT WHERE NAME LIKE 'FILE_%'")
for (const row of rows) {
    console.println(`${row.NAME}: ${row.VALUE}`);
}
rows.close();
conn.close();
db.close();
