/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
!function(){function r(r,e){var o=Object.keys(r);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(r);e&&(t=t.filter((function(e){return Object.getOwnPropertyDescriptor(r,e).enumerable}))),o.push.apply(o,t)}return o}function e(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?r(Object(a),!0).forEach((function(r){o(e,r,a[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):r(Object(a)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(a,r))}))}return e}function o(r,e,o){var t;return(e="symbol"==typeof(t=function(r,e){if("object"!=typeof r||!r)return r;var o=r[Symbol.toPrimitive];if(void 0!==o){var t=o.call(r,e||"default");if("object"!=typeof t)return t;throw new TypeError("@@toPrimitive must return a primitive value.")}return("string"===e?String:Number)(r)}(e,"string"))?t:t+"")in r?Object.defineProperty(r,e,{value:o,enumerable:!0,configurable:!0,writable:!0}):r[e]=o,r}System.register(["./087AC4D233B64EB0index-legacy.5acf65b9.js"],(function(r,o){"use strict";var t,a,d,n,l,i,c,s,b,u,f,p,v,y,g,h;return{setters:[function(r){t=r.B,a=r.ad,d=r.v,n=r.x,l=r.o,i=r.c,c=r.n,s=r.F,b=r.A,u=r.f,f=r.t,p=r.e,v=r.b,y=r.G,g=r.J,h=r.K}],execute:function(){var o=document.createElement("style");o.textContent=".el-card{--el-card-border-color: var(--el-border-color-light);--el-card-border-radius: 4px;--el-card-padding: 20px;--el-card-bg-color: var(--el-fill-color-blank);border-radius:var(--el-card-border-radius);border:1px solid var(--el-card-border-color);background-color:var(--el-card-bg-color);overflow:hidden;color:var(--el-text-color-primary);transition:var(--el-transition-duration)}.el-card.is-always-shadow{box-shadow:var(--el-box-shadow-light)}.el-card.is-hover-shadow:hover,.el-card.is-hover-shadow:focus{box-shadow:var(--el-box-shadow-light)}.el-card__header{padding:calc(var(--el-card-padding) - 2px) var(--el-card-padding);border-bottom:1px solid var(--el-card-border-color);box-sizing:border-box}.el-card__body{padding:var(--el-card-padding)}.el-card__footer{padding:calc(var(--el-card-padding) - 2px) var(--el-card-padding);border-top:1px solid var(--el-card-border-color);box-sizing:border-box}\n",document.head.appendChild(o);var w=t({header:{type:String,default:""},footer:{type:String,default:""},bodyStyle:{type:a([String,Object,Array]),default:""},bodyClass:String,shadow:{type:String,values:["always","hover","never"],default:"always"}}),x=d({name:"ElCard"}),O=d(e(e({},x),{},{props:w,setup:function(r){var e=n("card");return function(r,o){return l(),i("div",{class:c([s(e).b(),s(e).is("".concat(r.shadow,"-shadow"))])},[r.$slots.header||r.header?(l(),i("div",{key:0,class:c(s(e).e("header"))},[b(r.$slots,"header",{},(function(){return[u(f(r.header),1)]}))],2)):p("v-if",!0),v("div",{class:c([s(e).e("body"),r.bodyClass]),style:y(r.bodyStyle)},[b(r.$slots,"default")],6),r.$slots.footer||r.footer?(l(),i("div",{key:1,class:c(s(e).e("footer"))},[b(r.$slots,"footer",{},(function(){return[u(f(r.footer),1)]}))],2)):p("v-if",!0)],2)}}}));r("E",h(g(O,[["__file","card.vue"]])))}}}))}();
