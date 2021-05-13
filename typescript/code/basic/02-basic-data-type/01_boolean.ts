/**
 * 布尔类型学习，内容要点：
 * 1. 定义语法  let ok: boolean = true
 * 2. 关键字为 boolean，在Go中为bool
 * 3. 类型注解，相当于强类型语言中类型声明
 * 4. 类型注解语法： (变量/函数): 数据类型  let ok: boolean = true ,  function fn(): res{}
 */

 let ok: boolean = true

 if (ok){
     console.log("Hello  ", ok);
 }
 
 
 if (!ok){
     console.log("Hello ", ok);
 }else{
     console.log("Sorry  ",!ok);
 }