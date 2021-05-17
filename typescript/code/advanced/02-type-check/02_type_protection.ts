/**
 * 类型保护，内容要点：
 * 1. 类型保护是指缩小类型的范围，在一定的块级作用域内由编译器推导其类型，提示并规避不合法的操作
 *    TypeScript 能够在特定的区块中保证变量属于某种确定的类型，可在次区块中心中引用此类型的属性或调用其方法
 * 2. 关键字: typeof, instance, in
 * 3. 类型保护 TypeScript 类型检查机制的第二个部分，通过 typeof、instanceof、in 和 字面量类型 将代码分割成范围更小的代码块
 * 4. typeof 判断基础类型
 * 5. instaceof 判断是否为某个对象的实例
 * 6. in  操作符用于确定属性是否存在于某个对象上，这也是一种缩小范围的类型保护
 * 7. 字面类型保护 判断是否等于某个指定值，如果等于，就会缩小范围
 * 8. 类型保护函数，字面类型保护的变种
 */


// 1. 不使用类型保护机制，代码冗余
enum LangType { Strong, Weak }
class Go {
    go: string
    helloGo() { }
}

class TypeScript {
    typescript: string
    helloTypeScript() { }
}

function getLang(langType: LangType) {
    let lang = langType === LangType.Strong ? new Go() : new TypeScript()

    // 多处需要代码，代码冗余也不优雅
    if ((lang as Go).helloGo) {
        (lang as Go).helloGo()
    } else {
        (lang as TypeScript).helloTypeScript()
    }
}

// 2. 使用 instanceof
function getLangForInstanceof(langType: LangType) {
    let lang = langType === LangType.Strong ? new Go() : new TypeScript()

    // 判断条件修改为：lang instanceof Go
    if (lang instanceof Go) {
        lang.helloGo()
    } else {
        lang.helloTypeScript()
    }
}


// 3. 使用 in
function getLangForIn(langType: LangType) {
    let lang = langType === LangType.Strong ? new Go() : new TypeScript()

    // 判断条件修改为：'go' in Go
    if ('go' in lang) {
        lang.helloGo()
    } else {
        lang.helloTypeScript()
    }
}

// 4. 类型保护函数，字面量类型保护的变种
function isGo(lang: Go | TypeScript): lang is Go {
    return (lang as Go).helloGo !== undefined
}

function getLangForFunc(langType: LangType) {
    let lang = langType === LangType.Strong ? new Go() : new TypeScript()

    // 判断条件修改为：isGo(lang)
    if (isGo(lang)) {
        lang.helloGo()
    } else {
        lang.helloTypeScript()
    }
}


// 5. 使用字面量类型保护
type Success = {
    success: true,
    code: number,
    object: object
}

type Fail = {
    success: false,
    code: number,
    errMsg: string,
    request: string
}

function test(arg: Success | Fail) {
    // 通过布尔字面量，将这个代码块中变量 arg 的类型限定为 Success 类型
    if (arg.success === true) {
        console.log(arg.object) // OK
        // console.log(arg.errMsg) // Error, Property 'errMsg' does not exist on type 'Success'
    } else {
        console.log(arg.errMsg) // OK
        // console.log(arg.object) // Error, Property 'object' does not exist on type 'Fail'
    }
}