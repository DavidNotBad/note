```css
/* 定位属性的值 */
#flrs > a[href="/forexample/about.html"]
#J_login_form > dl > dt > input[id='J_password']
#flrs > input[name="username"][type="text"]

/* 同时匹配一个元素的多个class */
button.btn.btn_big.btn_submit

/* 匹配前缀 */
a[id^='id_prefix_']

/* 匹配后缀 */
a[id$='_id_sufix']

/* 属性中包含特定字符 */
a[id*='id_pattern']

/* 下一个兄弟姐妹 */
#username + input

/* 选择第n个特定元素 */
#recordlist li:nth-of-type(n)

/* 在父节点中选择第4个特定元素 */
#recordlist li:nth-child(n)

/* 在父节点中选择第n个元素, 不考虑类型 */
#recordlist *：nth-child(n)

/* 匹配文本 */
a:contains('Log Out')
```

