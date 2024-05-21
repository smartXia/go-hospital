/*
 *  web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 devops-manage
       当前版本:v2.6.4
    `)
  }
}
