/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{B as e,ad as a,aN as t,v as l,aa as s,aK as o,x as n,a as r,V as i,ak as u,Y as c,o as d,c as v,n as b,F as p,G as f,J as m,a2 as y,d6 as h,d7 as g,D as x,X as C,b8 as B,d as P,I as w,a6 as T,a9 as k,bL as N,aD as A,y as E,A as S,d3 as R,bm as F,ac as K,$ as L,T as D,U as _,d8 as V,r as j,ab as z,a4 as q,a7 as M,e as X,K as Y,L as $}from"./087AC4D233B64EB0index.5fcd1422.js";import{c as G}from"./087AC4D233B64EB0strings.84714f21.js";import{u as H}from"./087AC4D233B64EB0index.0de793c7.js";const I=Symbol("tabsRootContextKey"),J=e({tabs:{type:a(Array),default:()=>t([])}}),O="ElTabBar",U=l({name:O});var W=m(l({...U,props:J,setup(e,{expose:a}){const t=e,l=y(),m=s(I);m||o(O,"<el-tabs><el-tab-bar /></el-tabs>");const h=n("tabs"),g=r(),x=r(),C=()=>x.value=(()=>{let e=0,a=0;const s=["top","bottom"].includes(m.props.tabPosition)?"width":"height",o="width"===s?"x":"y",n="x"===o?"left":"top";return t.tabs.every((o=>{var r,i;const u=null==(i=null==(r=l.parent)?void 0:r.refs)?void 0:i["tab-".concat(o.uid)];if(!u)return!1;if(!o.active)return!0;e=u["offset".concat(G(n))],a=u["client".concat(G(s))];const c=window.getComputedStyle(u);return"width"===s&&(t.tabs.length>1&&(a-=Number.parseFloat(c.paddingLeft)+Number.parseFloat(c.paddingRight)),e+=Number.parseFloat(c.paddingLeft)),!1})),{[s]:"".concat(a,"px"),transform:"translate".concat(G(o),"(").concat(e,"px)")}})();return i((()=>t.tabs),(async()=>{await u(),C()}),{immediate:!0}),c(g,(()=>C())),a({ref:g,update:C}),(e,a)=>(d(),v("div",{ref_key:"barRef",ref:g,class:b([p(h).e("active-bar"),p(h).is(p(m).props.tabPosition)]),style:f(x.value)},null,6))}}),[["__file","tab-bar.vue"]]);const Q=e({panes:{type:a(Array),default:()=>t([])},currentName:{type:[String,Number],default:""},editable:Boolean,type:{type:String,values:["card","border-card",""],default:""},stretch:Boolean}),Z="ElTabNav",ee=l({name:Z,props:Q,emits:{tabClick:(e,a,t)=>t instanceof Event,tabRemove:(e,a)=>a instanceof Event},setup(e,{expose:a,emit:t}){const l=y(),d=s(I);d||o(Z,"<el-tabs><tab-nav /></el-tabs>");const v=n("tabs"),b=h(),p=g(),f=r(),m=r(),E=r(),S=r(),R=r(!1),F=r(0),K=r(!1),L=r(!0),D=x((()=>["top","bottom"].includes(d.props.tabPosition)?"width":"height")),_=x((()=>{const e="width"===D.value?"X":"Y";return{transform:"translate".concat(e,"(-").concat(F.value,"px)")}})),V=()=>{if(!f.value)return;const e=f.value["offset".concat(G(D.value))],a=F.value;if(!a)return;const t=a>e?a-e:0;F.value=t},j=()=>{if(!f.value||!m.value)return;const e=m.value["offset".concat(G(D.value))],a=f.value["offset".concat(G(D.value))],t=F.value;if(e-t<=a)return;const l=e-t>2*a?t+a:e-a;F.value=l},z=async()=>{const e=m.value;if(!(R.value&&E.value&&f.value&&e))return;await u();const a=E.value.querySelector(".is-active");if(!a)return;const t=f.value,l=["top","bottom"].includes(d.props.tabPosition),s=a.getBoundingClientRect(),o=t.getBoundingClientRect(),n=l?e.offsetWidth-o.width:e.offsetHeight-o.height,r=F.value;let i=r;l?(s.left<o.left&&(i=r-(o.left-s.left)),s.right>o.right&&(i=r+s.right-o.right)):(s.top<o.top&&(i=r-(o.top-s.top)),s.bottom>o.bottom&&(i=r+(s.bottom-o.bottom))),i=Math.max(i,0),F.value=Math.min(i,n)},q=()=>{var a;if(!m.value||!f.value)return;e.stretch&&(null==(a=S.value)||a.update());const t=m.value["offset".concat(G(D.value))],l=f.value["offset".concat(G(D.value))],s=F.value;l<t?(R.value=R.value||{},R.value.prev=s,R.value.next=s+l<t,t-s<l&&(F.value=t-l)):(R.value=!1,s>0&&(F.value=0))},M=e=>{const a=e.code,{up:t,down:l,left:s,right:o}=A;if(![t,l,s,o].includes(a))return;const n=Array.from(e.currentTarget.querySelectorAll("[role=tab]:not(.is-disabled)")),r=n.indexOf(e.target);let i;i=a===s||a===t?0===r?n.length-1:r-1:r<n.length-1?r+1:0,n[i].focus({preventScroll:!0}),n[i].click(),X()},X=()=>{L.value&&(K.value=!0)},Y=()=>K.value=!1;return i(b,(e=>{"hidden"===e?L.value=!1:"visible"===e&&setTimeout((()=>L.value=!0),50)})),i(p,(e=>{e?setTimeout((()=>L.value=!0),50):L.value=!1})),c(E,q),C((()=>setTimeout((()=>z()),0))),B((()=>q())),a({scrollToActiveTab:z,removeFocus:Y}),i((()=>e.panes),(()=>l.update()),{flush:"post",deep:!0}),()=>{const a=R.value?[P("span",{class:[v.e("nav-prev"),v.is("disabled",!R.value.prev)],onClick:V},[P(w,null,{default:()=>[P(T,null,null)]})]),P("span",{class:[v.e("nav-next"),v.is("disabled",!R.value.next)],onClick:j},[P(w,null,{default:()=>[P(k,null,null)]})])]:null,l=e.panes.map(((a,l)=>{var s,o,n,r;const i=a.uid,u=a.props.disabled,c=null!=(o=null!=(s=a.props.name)?s:a.index)?o:"".concat(l),b=!u&&(a.isClosable||e.editable);a.index="".concat(l);const p=b?P(w,{class:"is-icon-close",onClick:e=>t("tabRemove",a,e)},{default:()=>[P(N,null,null)]}):null,f=(null==(r=(n=a.slots).label)?void 0:r.call(n))||a.props.label,m=!u&&a.active?0:-1;return P("div",{ref:"tab-".concat(i),class:[v.e("item"),v.is(d.props.tabPosition),v.is("active",a.active),v.is("disabled",u),v.is("closable",b),v.is("focus",K.value)],id:"tab-".concat(c),key:"tab-".concat(i),"aria-controls":"pane-".concat(c),role:"tab","aria-selected":a.active,tabindex:m,onFocus:()=>X(),onBlur:()=>Y(),onClick:e=>{Y(),t("tabClick",a,c,e)},onKeydown:e=>{!b||e.code!==A.delete&&e.code!==A.backspace||t("tabRemove",a,e)}},[f,p])}));return P("div",{ref:E,class:[v.e("nav-wrap"),v.is("scrollable",!!R.value),v.is(d.props.tabPosition)]},[a,P("div",{class:v.e("nav-scroll"),ref:f},[P("div",{class:[v.e("nav"),v.is(d.props.tabPosition),v.is("stretch",e.stretch&&["top","bottom"].includes(d.props.tabPosition))],ref:m,style:_.value,role:"tablist",onKeydown:M},[e.type?null:P(W,{ref:S,tabs:[...e.panes]},null),l])])])}}}),ae=e({type:{type:String,values:["card","border-card",""],default:""},closable:Boolean,addable:Boolean,modelValue:{type:[String,Number]},editable:Boolean,tabPosition:{type:String,values:["top","right","bottom","left"],default:"top"},beforeLeave:{type:a(Function),default:()=>!0},stretch:Boolean}),te=e=>L(e)||D(e),le=l({name:"ElTabs",props:ae,emits:{[F]:e=>te(e),tabClick:(e,a)=>a instanceof Event,tabChange:e=>te(e),edit:(e,a)=>["remove","add"].includes(a),tabRemove:e=>te(e),tabAdd:()=>!0},setup(e,{emit:a,slots:t,expose:l}){var s;const o=n("tabs"),{children:c,addChild:d,removeChild:v}=H(y(),"ElTabPane"),b=r(),p=r(null!=(s=e.modelValue)?s:"0"),f=async(t,l=!1)=>{var s,o,n;if(p.value!==t&&!K(t))try{!1!==await(null==(s=e.beforeLeave)?void 0:s.call(e,t,p.value))&&(p.value=t,l&&(a(F,t),a("tabChange",t)),null==(n=null==(o=b.value)?void 0:o.removeFocus)||n.call(o))}catch(r){}},m=(e,t,l)=>{e.props.disabled||(f(t,!0),a("tabClick",e,l))},h=(e,t)=>{e.props.disabled||K(e.props.name)||(t.stopPropagation(),a("edit",e.props.name,"remove"),a("tabRemove",e.props.name))},g=()=>{a("edit",void 0,"add"),a("tabAdd")};return i((()=>e.modelValue),(e=>f(e))),i(p,(async()=>{var e;await u(),null==(e=b.value)||e.scrollToActiveTab()})),E(I,{props:e,currentName:p,registerPane:d,unregisterPane:v}),l({currentName:p}),()=>{const a=t["add-icon"],l=e.editable||e.addable?P("span",{class:o.e("new-tab"),tabindex:"0",onClick:g,onKeydown:e=>{e.code===A.enter&&g()}},[a?S(t,"add-icon"):P(w,{class:o.is("icon-plus")},{default:()=>[P(R,null,null)]})]):null,s=P("div",{class:[o.e("header"),o.is(e.tabPosition)]},[l,P(ee,{ref:b,currentName:p.value,editable:e.editable,type:e.type,panes:c.value,stretch:e.stretch,onTabClick:m,onTabRemove:h},null)]),n=P("div",{class:o.e("content")},[S(t,"default")]);return P("div",{class:[o.b(),o.m(e.tabPosition),{[o.m("card")]:"card"===e.type,[o.m("border-card")]:"border-card"===e.type}]},[..."bottom"!==e.tabPosition?[s,n]:[n,s]])}}}),se=e({label:{type:String,default:""},name:{type:[String,Number]},closable:Boolean,disabled:Boolean,lazy:Boolean}),oe=["id","aria-hidden","aria-labelledby"],ne="ElTabPane",re=l({name:ne});var ie=m(l({...re,props:se,setup(e){const a=e,t=y(),l=_(),u=s(I);u||o(ne,"usage: <el-tabs><el-tab-pane /></el-tabs/>");const c=n("tab-pane"),f=r(),m=x((()=>a.closable||u.props.closable)),h=V((()=>{var e;return u.currentName.value===(null!=(e=a.name)?e:f.value)})),g=r(h.value),B=x((()=>{var e;return null!=(e=a.name)?e:f.value})),P=V((()=>!a.lazy||g.value||h.value));i(h,(e=>{e&&(g.value=!0)}));const w=j({uid:t.uid,slots:l,props:a,paneName:B,active:h,index:f,isClosable:m});return C((()=>{u.registerPane(w)})),z((()=>{u.unregisterPane(w.uid)})),(e,a)=>p(P)?q((d(),v("div",{key:0,id:"pane-".concat(p(B)),class:b(p(c).b()),role:"tabpanel","aria-hidden":!p(h),"aria-labelledby":"tab-".concat(p(B))},[S(e.$slots,"default")],10,oe)),[[M,p(h)]]):X("v-if",!0)}}),[["__file","tab-pane.vue"]]);const ue=Y(le,{TabPane:ie}),ce=$(ie);export{ce as E,ue as a};
