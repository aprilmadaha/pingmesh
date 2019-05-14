$(function(){
    var width = String((parseInt($('h5').html())+1)*70)+'px';
    $('body').css('width',width);
    $('body').css('height',width);
    $('#content').css('width',width);
    function getColor(bili){
        //var 百分之一 = (单色值范围) / 50;  单颜色的变化范围只在50%之内
        var one = (255+255) / 100;  
        var r=0;
        var g=0;
        var b=0;

        if ( bili < 50 ) { 
            // 比例小于50的时候红色是越来越多的,直到红色为255时(红+绿)变为黄色.
            r = one * bili;
            g=255;
        }
        if ( bili >= 50 ) {
            // 比例大于50的时候绿色是越来越少的,直到0 变为纯红
            g =  255 - ( (bili - 50 ) * one) ;
            r = 255;
        }
        r = parseInt(r);// 取整
        g = parseInt(g);// 取整
        b = parseInt(b);// 取整
        return "rgb("+r+","+g+","+b+")";
            
      };
    // 等待css渲染
    setTimeout(interval_update, 5000);
    update_mesh();
    // 每三秒更新
    function interval_update(){
        setInterval(update_mesh, 30000);
    }
    
    function update_mesh(){
        $.get(
            '/update_mesh/',
            {},
            function(res){
                for (var k in res){
                    // console.log(k, res[k]);
                    $('#content').children('div[name="'+k+'"]').html(res[k][0]).attr('title',res[k][1]).css({'background-color':getColor(res[k][0])});
                }
            })
    }

})