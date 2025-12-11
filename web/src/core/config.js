/**
 * 网站配置文件
 */

const config = {
  appName: 'Hertz-Vue-Admin',
  appLogo: '@/assets/nav_logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    // chalk 为 ESM，仅在支持 import 的环境下使用
    import('chalk').then(({ default: chalk }) => {
      console.log(
        chalk.green(
          `> 欢迎使用Hertz-Vue-Admin，开源地址：https://github.com/EduFriendChen/hertz-vue-admin`
        )
      )
      console.log(
        chalk.green(
          `> 当前版本: 先行版`
        )
      )
      console.log('\n')
    }).catch(() => {
      console.log('> 欢迎使用Hertz-Vue-Admin，开源地址：https://github.com/EduFriendChen/hertz-vue-admin')
      console.log('> 当前版本: 先行版')
      console.log('\n')
    })
  }
}

export default config
