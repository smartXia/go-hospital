/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{af as e,l as a,u as l,au as r,a as o,r as t,V as s,o as n,g as d,w as u,a4 as i,b as c,b1 as m,d as p,f,c as h,N as b,O as B,t as g,e as k,E as v,P as y,n as C,F as w,az as E,I as D}from"./087AC4D233B64EB0index.5fcd1422.js";/* empty css                                *//* empty css                               */import j from"./087AC4D233B64EB0index.dac3bac5.js";import{E as A}from"./087AC4D233B64EB0dialog.e17c90b5.js";import"./087AC4D233B64EB0overlay.1520877c.js";/* empty css                               */import{E as _}from"./087AC4D233B64EB0index.71b3c6ca.js";import"./087AC4D233B64EB0drawer.6a894563.js";/* empty css                            *//* empty css                              */import"./087AC4D233B64EB0input-number.f5afd020.js";import"./087AC4D233B64EB0index.c25afd12.js";import"./087AC4D233B64EB0switch.5175e219.js";import"./087AC4D233B64EB0position.18b59761.js";import"./087AC4D233B64EB0index.508774e7.js";import"./087AC4D233B64EB0debounce.679d5750.js";import"./087AC4D233B64EB0index.d064f019.js";const x={key:0,class:"quick-title"},O=["onClick"],T={class:"dialog-footer"},q=Object.assign({name:"CommandMenu"},{__name:"index",setup(y,{expose:C}){const w=e(),E=a(),D=l(),j=l(),_=r(),q=o(!1),M=o(""),V=t([]),S=e=>{const a=[];return e.forEach((e=>{e.children&&e.children.length>0?a.push(...S(e.children)):e.meta.title&&e.meta.title.indexOf(M.value)>-1&&a.push({label:e.meta.title,func:()=>U(e)})})),a},I=()=>{const e={label:"跳转",children:[]},a=S(_.asyncRouters[0].children);e.children.push(...a),V.push(e)},L=()=>{const e={label:"操作",children:[]},a=[{label:"亮色主题",func:()=>N("light")},{label:"暗色主题",func:()=>N("dark")},{label:"退出登录",func:()=>E.LoginOut()}];e.children.push(...a.filter((e=>e.label.indexOf(M.value)>-1))),V.push(e)};I(),L();const U=e=>{var a,l;const r=e.name,o={},t={};(null==(a=_.routeMap[r])?void 0:a.parameters)&&(null==(l=_.routeMap[r])||l.parameters.forEach((e=>{"query"===e.type?o[e.key]=e.value:t[e.key]=e.value}))),r!==j.name&&(e.name.indexOf("http://")>-1||e.name.indexOf("https://")>-1?window.open(e.name):D.push({name:r,query:o,params:t}),q.value=!1)},N=e=>{null!==e?w.toggleTheme(!0):w.toggleTheme(!1)},R=()=>{q.value=!1};return C({open:()=>{q.value=!0}}),s(M,(()=>{V.length=0,I(),L()})),(e,a)=>{const l=v,r=A;return n(),d(r,{modelValue:q.value,"onUpdate:modelValue":a[1]||(a[1]=e=>q.value=e),width:"30%",class:"overlay","show-close":!1},{header:u((()=>[i(c("input",{"onUpdate:modelValue":a[0]||(a[0]=e=>M.value=e),class:"quick-input",placeholder:"请输入你需要快捷到达的功能"},null,512),[[m,M.value]])])),footer:u((()=>[c("span",T,[p(l,{onClick:R},{default:u((()=>[f("关闭")])),_:1})])])),default:u((()=>[(n(!0),h(b,null,B(V,((e,a)=>(n(),h("div",{key:a},[e.children.length?(n(),h("div",x,g(e.label),1)):k("",!0),(n(!0),h(b,null,B(e.children,((e,l)=>(n(),h("div",{key:a+"-"+l,class:"quick-item",onClick:e.func},g(e.label),9,O)))),128))])))),128))])),_:1},8,["modelValue"])}}}),M={class:"flex items-center mx-8 gap-4"},V={__name:"tools",setup(a){const l=e(),r=o(!1),t=o(!1),s=()=>{t.value=!0,E.emit("reload"),setTimeout((()=>{t.value=!1}),1e3)},i=()=>{r.value=!0},c=o(""),m=o(),f=()=>{m.value.open()};return(()=>{"WIN"===window.localStorage.getItem("osType")?c.value="Ctrl":c.value="⌘";window.addEventListener("keydown",(e=>{e.ctrlKey&&"k"===e.key&&(e.preventDefault(),f())}))})(),(e,a)=>{const o=y("Search"),c=D,b=_,B=y("Setting"),g=y("Refresh"),k=y("Sunny"),v=y("Moon");return n(),h("div",M,[p(b,{class:"",effect:"dark",content:"搜索",placement:"bottom"},{default:u((()=>[p(c,{onClick:f,class:"w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid"},{default:u((()=>[p(o)])),_:1})])),_:1}),p(b,{class:"",effect:"dark",content:"系统设置",placement:"bottom"},{default:u((()=>[p(c,{class:"w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid",onClick:i},{default:u((()=>[p(B)])),_:1})])),_:1}),p(b,{class:"",effect:"dark",content:"刷新",placement:"bottom"},{default:u((()=>[p(c,{class:C(["w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid",t.value?"animate-spin":""]),onClick:s},{default:u((()=>[p(g)])),_:1},8,["class"])])),_:1}),p(b,{class:"",effect:"dark",content:"切换主题",placement:"bottom",disabled:"auto"===w(l).theme},{default:u((()=>["dark"===w(l).theme?(n(),d(c,{key:0,class:"w-8 h-8 shadow rounded-full border border-gray-600 cursor-pointer border-solid",onClick:a[0]||(a[0]=e=>w(l).toggleTheme(!1))},{default:u((()=>[p(k)])),_:1})):(n(),d(c,{key:1,class:"w-8 h-8 shadow rounded-full border border-gray-200 cursor-pointer border-solid",onClick:a[1]||(a[1]=e=>w(l).toggleTheme(!0))},{default:u((()=>[p(v)])),_:1}))])),_:1},8,["disabled"]),p(j,{drawer:r.value,"onUpdate:drawer":a[2]||(a[2]=e=>r.value=e)},null,8,["drawer"]),p(q,{ref_key:"command",ref:m},null,512)])}}};export{V as default};
