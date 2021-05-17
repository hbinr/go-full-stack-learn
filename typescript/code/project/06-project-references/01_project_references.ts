/**
 *
 * 1.
 * 2. 新增 composite 配置属性，为true，表示工程之间可以引用
 *    提高了工程性，每个功能模块都可以有自己的tsconfig.json
 * 3. 新增 declaration 配置属性，为true，表示编译后会生成声明文件，这是工程引用所必须的
 * 4. 构建命令:  tsc -b src/client --verbose, 支持单工程构建
 *      a. -b  -build 简写
 *      b. --verbose  输出构建信息
 *      c. --clean  清空构建的文件  eg: tsc -b test --verbose
 *
 * 5. 工程引用优点：
 *      a. 解决了输出目录的结构问题
 *      b. 解决了单个工程构建的问题
 *      c. 通过增量编译，提升了构建速度
 *
 */