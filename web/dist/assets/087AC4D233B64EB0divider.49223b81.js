/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{B as e,ad as t,v as s,x as a,D as i,o as r,c as o,n as l,F as n,A as d,e as c,G as v,J as u,K as f}from"./087AC4D233B64EB0index.5fcd1422.js";const p=e({direction:{type:String,values:["horizontal","vertical"],default:"horizontal"},contentPosition:{type:String,values:["left","center","right"],default:"center"},borderStyle:{type:t(String),default:"solid"}}),y=s({name:"ElDivider"});const S=f(u(s({...y,props:p,setup(e){const t=e,s=a("divider"),u=i((()=>s.cssVar({"border-style":t.borderStyle})));return(e,t)=>(r(),o("div",{class:l([n(s).b(),n(s).m(e.direction)]),style:v(n(u)),role:"separator"},[e.$slots.default&&"vertical"!==e.direction?(r(),o("div",{key:0,class:l([n(s).e("text"),n(s).is(e.contentPosition)])},[d(e.$slots,"default")],2)):c("v-if",!0)],6))}}),[["__file","divider.vue"]]));export{S as E};
