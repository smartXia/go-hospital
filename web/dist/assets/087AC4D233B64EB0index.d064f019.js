/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{B as s,b2 as a,v as e,aY as o,x as l,D as n,o as t,c,b as i,A as r,n as u,F as p,g as d,w as f,d as k,bL as m,a5 as g,I as y,e as b,G as v,a8 as C,J as B,K as h}from"./087AC4D233B64EB0index.5fcd1422.js";const E=s({type:{type:String,values:["primary","success","info","warning","danger"],default:"primary"},closable:Boolean,disableTransitions:Boolean,hit:Boolean,color:String,size:{type:String,values:a},effect:{type:String,values:["dark","light","plain"],default:"light"},round:Boolean}),_={close:s=>s instanceof MouseEvent,click:s=>s instanceof MouseEvent},S=e({name:"ElTag"});const x=h(B(e({...S,props:E,emits:_,setup(s,{emit:a}){const e=s,B=o(),h=l("tag"),E=n((()=>{const{type:s,hit:a,effect:o,closable:l,round:n}=e;return[h.b(),h.is("closable",l),h.m(s||"primary"),h.m(B.value),h.m(o),h.is("hit",a),h.is("round",n)]})),_=s=>{a("close",s)},S=s=>{a("click",s)};return(s,a)=>s.disableTransitions?(t(),c("span",{key:0,class:u(p(E)),style:v({backgroundColor:s.color}),onClick:S},[i("span",{class:u(p(h).e("content"))},[r(s.$slots,"default")],2),s.closable?(t(),d(p(y),{key:0,class:u(p(h).e("close")),onClick:g(_,["stop"])},{default:f((()=>[k(p(m))])),_:1},8,["class","onClick"])):b("v-if",!0)],6)):(t(),d(C,{key:1,name:"".concat(p(h).namespace.value,"-zoom-in-center"),appear:""},{default:f((()=>[i("span",{class:u(p(E)),style:v({backgroundColor:s.color}),onClick:S},[i("span",{class:u(p(h).e("content"))},[r(s.$slots,"default")],2),s.closable?(t(),d(p(y),{key:0,class:u(p(h).e("close")),onClick:g(_,["stop"])},{default:f((()=>[k(p(m))])),_:1},8,["class","onClick"])):b("v-if",!0)],6)])),_:3},8,["name"]))}}),[["__file","tag.vue"]]));export{x as E,E as t};
