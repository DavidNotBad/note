## 求和

```php
$orders = [
    [
        'a' =>  'a1',
        'b' =>  [
            ['c'=>33],
            ['c'=>11]
        ]
    ]
];

collect($orders)->map(function($order){
    return $order['b'];
})->flatten(1)->map(function($order){
    return $order['c'];
})->sum();

collect($orders)->flatMap(function($order){
    return $order['b'];
})->pluck('c')->sum();

collect($orders)->flatMap(function($order){
    return $order['b'];
})->sum('c')
    
collect($orders)->pluck('b.*.c')->flatten()->sum();
```

