/**
 * 网站配置文件
 */
import chalk from "chalk";
const config = {
  appName: "Gin-Vue-JNY",
  appLogo: "https://s2-cdn.oneitfarm.com/e12ece12233b4059860acf48637d830a.jpg",
  showViteLogo: true,
  logs: [],
};

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
   
    console.log(chalk.green(`** some happy day  **`));
    console.log('\n')
  }
}

export default config
