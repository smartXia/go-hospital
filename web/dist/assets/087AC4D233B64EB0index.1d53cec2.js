/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{_ as e,S as a,u as s,a as t,l,D as n,V as u,az as r,ab as o,o as i,c as v,d as m,w as c,N as d,O as p,g as y,b as g,n as f,f as h,t as b,F as S,a$ as x,a5 as I,a4 as q,a7 as w,G as k,ak as O}from"./087AC4D233B64EB0index.5fcd1422.js";import{E as C,a as N}from"./087AC4D233B64EB0tab-pane.20a0e58d.js";import"./087AC4D233B64EB0strings.84714f21.js";import"./087AC4D233B64EB0index.0de793c7.js";const E={class:"gva-tabs"},J=["tab"],j=e(Object.assign({name:"HistoryComponent"},{__name:"index",setup(e){const j=a(),A=s(),B=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),V=t([]),_=t(""),D=t(!1),P=l(),T=t(0),L=t(0),R=t(!1),z=t(!1),U=t(""),$=n((()=>P.userInfo.authority.defaultRouter)),F=()=>{V.value=[{name:$.value,meta:{title:"首页"},query:{},params:{}}],A.push({name:$.value}),D.value=!1,sessionStorage.setItem("historys",JSON.stringify(V.value))},G=()=>{let e;const a=V.value.findIndex((a=>(B(a)===U.value&&(e=a),B(a)===U.value))),s=V.value.findIndex((e=>B(e)===_.value));V.value.splice(0,a),a>s&&A.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},H=()=>{let e;const a=V.value.findIndex((a=>(B(a)===U.value&&(e=a),B(a)===U.value))),s=V.value.findIndex((e=>B(e)===_.value));V.value.splice(a+1,V.value.length),a<s&&A.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},K=()=>{let e;V.value=V.value.filter((a=>(B(a)===U.value&&(e=a),B(a)===U.value))),A.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},Q=e=>{if(!V.value.some((a=>((e,a)=>{if(e.name!==a.name)return!1;if(Object.keys(e.query).length!==Object.keys(a.query).length||Object.keys(e.params).length!==Object.keys(a.params).length)return!1;for(const s in e.query)if(e.query[s]!==a.query[s])return!1;for(const s in e.params)if(e.params[s]!==a.params[s])return!1;return!0})(a,e)))){const a={};a.name=e.name,a.meta={...e.meta},delete a.meta.matched,a.query=e.query,a.params=e.params,V.value.push(a)}window.sessionStorage.setItem("activeValue",B(e))},X=t({}),Y=e=>{var a;const s=null==(a=null==e?void 0:e.props)?void 0:a.name;if(!s)return;const t=X.value[s];A.push({name:t.name,query:t.query,params:t.params})},M=e=>{const a=V.value.findIndex((a=>B(a)===e));B(j)===e&&(1===V.value.length?A.push({name:$.value}):a<V.value.length-1?A.push({name:V.value[a+1].name,query:V.value[a+1].query,params:V.value[a+1].params}):A.push({name:V.value[a-1].name,query:V.value[a-1].query,params:V.value[a-1].params})),V.value.splice(a,1)};u((()=>D.value),(()=>{D.value?document.body.addEventListener("click",(()=>{D.value=!1})):document.body.removeEventListener("click",(()=>{D.value=!1}))})),u((()=>j),((e,a)=>{"Login"!==e.name&&"Reload"!==e.name&&(V.value=V.value.filter((e=>!e.meta.closeTab)),Q(e),sessionStorage.setItem("historys",JSON.stringify(V.value)),_.value=window.sessionStorage.getItem("activeValue"))}),{deep:!0}),u((()=>V.value),(()=>{sessionStorage.setItem("historys",JSON.stringify(V.value)),X.value={},V.value.forEach((e=>{X.value[B(e)]=e})),r.emit("setKeepAlive",V.value)}),{deep:!0});return(()=>{r.on("closeThisPage",(()=>{M(B(j))})),r.on("closeAllPage",(()=>{F()})),r.on("mobile",(e=>{z.value=e})),r.on("collapse",(e=>{R.value=e})),r.on("setQuery",(e=>{const a=V.value.findIndex((e=>B(e)===_.value));V.value[a].query=e,_.value=B(V.value[a]);const s=window.location.href.split("?")[0],t=new URLSearchParams(e).toString();window.history.replaceState({},"","".concat(s,"?").concat(t)),sessionStorage.setItem("historys",JSON.stringify(V.value))})),r.on("switchTab",(async e=>{const a=V.value.findIndex((a=>a.name===e.name));if(!(a<0)){for(const a in e.query)e.query[a]=String(e.query[a]);for(const a in e.params)e.params[a]=String(e.params[a]);V.value[a].query=e.query||{},V.value[a].params=e.params||{},await O(),A.push(V.value[a])}}));const e=[{name:$.value,meta:{title:"首页"},query:{},params:{}}];Q(j),V.value=JSON.parse(sessionStorage.getItem("historys"))||e,window.sessionStorage.getItem("activeValue")?_.value=window.sessionStorage.getItem("activeValue"):_.value=B(j),"true"===window.sessionStorage.getItem("needCloseAll")&&(F(),window.sessionStorage.removeItem("needCloseAll"))})(),o((()=>{r.off("collapse"),r.off("mobile")})),(e,a)=>{const s=C,t=N;return i(),v("div",E,[m(t,{modelValue:_.value,"onUpdate:modelValue":a[0]||(a[0]=e=>_.value=e),closable:!(1===V.value.length&&e.$route.name===$.value),type:"card",class:"bg-white text-slate-700 dark:text-slate-500 dark:bg-slate-900",onContextmenu:a[1]||(a[1]=I((e=>(e=>{if(1===V.value.length&&j.name===$.value)return!1;let a="";a="SPAN"===e.srcElement.nodeName?e.srcElement.offsetParent.id:e.srcElement.id,a&&(D.value=!0,T.value=e.clientX,L.value=e.clientY+10,U.value=a.substring(4))})(e)),["prevent"])),onTabClick:Y,onTabRemove:M},{default:c((()=>[(i(!0),v(d,null,p(V.value,(e=>(i(),y(s,{key:B(e),label:e.meta.title,name:B(e),tab:e,class:"border-none"},{label:c((()=>[g("span",{tab:e,class:f(_.value===B(e)?"text-active":"text-gray-600 dark:text-slate-400 ")},[g("i",{class:f(_.value===B(e)?"text-active":"text-gray-600 dark:text-slate-400")},null,2),h(" "+b(S(x)(e.meta.title,e)),1)],10,J)])),_:2},1032,["label","name","tab"])))),128))])),_:1},8,["modelValue","closable"]),q(g("ul",{style:k({left:T.value+"px",top:L.value+"px"}),class:"contextmenu"},[g("li",{onClick:F}," 关闭所有 "),g("li",{onClick:G}," 关闭左侧 "),g("li",{onClick:H}," 关闭右侧 "),g("li",{onClick:K}," 关闭其他 ")],4),[[w,D.value]])])}}}),[["__scopeId","data-v-9508dab6"]]);export{j as default};
