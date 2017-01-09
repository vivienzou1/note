<?php
include('./simple_html_dom.php');

$str = getUrlInfo('http://m.ipin.com/major/576ce15ac62e9c35de92c84d-zk.html');
$html = str_get_html($str);

//专业介绍
foreach($html->find('div.job-intro-details div div') as $e){
    //echo trim($e->innertext) . '<br>';
}

//毕业后拿多少钱
foreach($html->find('div#salary') as $e){
    //echo trim($e->innertext) . '<br>';
}

//在什么行业
foreach($html->find('div#industry') as $e){
    //echo trim($e->innertext) . '<br>';
}

//行业特点
foreach($html->find('div.card') as $e){
    //echo trim($e->innertext) . '<br>';
}

//在哪里就业
foreach($html->find('div#area') as $e){
    //echo trim($e->innertext) . '<br>';
}

//地区特点
foreach($html->find('div.card p.forjob') as $e){
    //echo trim($e->innertext) . '<br>';
}

//哪些大学有这个专业
foreach($html->find('div#school') as $e){
    //echo trim($e->innertext) . '<br>';
}

//还有什么类似专业
foreach($html->find('div.padd-b0') as $e){
    echo trim($e->innertext) . '<br>';
}

function getUrlInfo($url)
{
    // 创建一个新cURL资源
    $ch = curl_init($url);

    // 设置cURL传输选项
    // 设置报头信息是否包含在数据流输出中
    curl_setopt($ch, CURLOPT_HEADER, FALSE);
    // 在HTTP请求中包含一个"User-Agent: "头的字符串
    curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36');
    // curl_exec() 将以字符串形式返回获取的信息，而不是在脚本执行时直接输出
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);

    // 根据服务器返回 HTTP 头中的 "Location: " 重定向(这是递归的，"Location: " 发送几次就重定向几次)
    // http://www.le.com/player/x1380568.swf
    // 301重定向=>
    // http://img1.c0.letv.com/ptv/player/swfPlayer.swf?id=1380568
    curl_setopt($ch, CURLOPT_FOLLOWLOCATION, TRUE);

    // 抓取URL并把返回的数据传递给$result
    $result = curl_exec($ch);
    // 关闭cURL资源，并且释放系统资源
    curl_close($ch);

    return $result;
}
