/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{B as e,b4 as t,v as a,D as l,x as o,a as r,F as s,aX as i,o as n,g as p,w as d,c as u,n as c,t as f,e as v,A as m,f as b,aF as g,J as h,dF as B,K as y,s as w,P as C,b as A,d as E,Q as _,R as D,j,k as x,E as k,I as S}from"./087AC4D233B64EB0index.5fcd1422.js";import{E as O}from"./087AC4D233B64EB0pagination.81705c9f.js";/* empty css                              *//* empty css                            */import"./087AC4D233B64EB0select.591b7648.js";import"./087AC4D233B64EB0scrollbar.b45aa0d9.js";/* empty css                               */import{E as R,a as N}from"./087AC4D233B64EB0table-column.de1f83f6.js";import"./087AC4D233B64EB0checkbox.2864820c.js";/* empty css                                */import{a as V,E as I}from"./087AC4D233B64EB0form-item.69b26929.js";/* empty css                               */import{E as U}from"./087AC4D233B64EB0index.d064f019.js";import{b as z,u as T,E as F}from"./087AC4D233B64EB0index.71b3c6ca.js";import{d as $}from"./087AC4D233B64EB0dropdown.25dd4390.js";import"./087AC4D233B64EB0isEqual.c192b418.js";import"./087AC4D233B64EB0_Uint8Array.72188c2c.js";import"./087AC4D233B64EB0strings.84714f21.js";import"./087AC4D233B64EB0debounce.679d5750.js";import"./087AC4D233B64EB0_baseIteratee.bf25bc12.js";import"./087AC4D233B64EB0index.508774e7.js";import"./087AC4D233B64EB0castArray.ef7e3470.js";import"./087AC4D233B64EB0_initCloneObject.67d0c1fc.js";import"./087AC4D233B64EB0_baseClone.bd09a709.js";const P=e({trigger:z.trigger,placement:$.placement,disabled:z.disabled,visible:T.visible,transition:T.transition,popperOptions:$.popperOptions,tabindex:$.tabindex,content:T.content,popperStyle:T.popperStyle,popperClass:T.popperClass,enterable:{...T.enterable,default:!0},effect:{...T.effect,default:"light"},teleported:T.teleported,title:String,width:{type:[String,Number],default:150},offset:{type:Number,default:void 0},showAfter:{type:Number,default:0},hideAfter:{type:Number,default:200},autoClose:{type:Number,default:0},showArrow:{type:Boolean,default:!0},persistent:{type:Boolean,default:!0},"onUpdate:visible":{type:Function}}),q={"update:visible":e=>t(e),"before-enter":()=>!0,"before-leave":()=>!0,"after-enter":()=>!0,"after-leave":()=>!0},H=a({name:"ElPopover"}),J=a({...H,props:P,emits:q,setup(e,{expose:t,emit:a}){const h=e,B=l((()=>h["onUpdate:visible"])),y=o("popover"),w=r(),C=l((()=>{var e;return null==(e=s(w))?void 0:e.popperRef})),A=l((()=>[{width:i(h.width)},h.popperStyle])),E=l((()=>[y.b(),h.popperClass,{[y.m("plain")]:!!h.content}])),_=l((()=>h.transition==="".concat(y.namespace.value,"-fade-in-linear"))),D=()=>{a("before-enter")},j=()=>{a("before-leave")},x=()=>{a("after-enter")},k=()=>{a("update:visible",!1),a("after-leave")};return t({popperRef:C,hide:()=>{var e;null==(e=w.value)||e.hide()}}),(e,t)=>(n(),p(s(F),g({ref_key:"tooltipRef",ref:w},e.$attrs,{trigger:e.trigger,placement:e.placement,disabled:e.disabled,visible:e.visible,transition:e.transition,"popper-options":e.popperOptions,tabindex:e.tabindex,content:e.content,offset:e.offset,"show-after":e.showAfter,"hide-after":e.hideAfter,"auto-close":e.autoClose,"show-arrow":e.showArrow,"aria-label":e.title,effect:e.effect,enterable:e.enterable,"popper-class":s(E),"popper-style":s(A),teleported:e.teleported,persistent:e.persistent,"gpu-acceleration":s(_),"onUpdate:visible":s(B),onBeforeShow:D,onBeforeHide:j,onShow:x,onHide:k}),{content:d((()=>[e.title?(n(),u("div",{key:0,class:c(s(y).e("title")),role:"title"},f(e.title),3)):v("v-if",!0),m(e.$slots,"default",{},(()=>[b(f(e.content),1)]))])),default:d((()=>[e.$slots.reference?m(e.$slots,"reference",{key:0}):v("v-if",!0)])),_:3},16,["trigger","placement","disabled","visible","transition","popper-options","tabindex","content","offset","show-after","hide-after","auto-close","show-arrow","aria-label","effect","enterable","popper-class","popper-style","teleported","persistent","gpu-acceleration","onUpdate:visible"]))}});const K=(e,t)=>{const a=t.arg||t.value,l=null==a?void 0:a.popperRef;l&&(l.triggerRef=e)};const L=y(h(J,[["__file","popover.vue"]]),{directive:B({mounted(e,t){K(e,t)},updated(e,t){K(e,t)}},"popover")}),Q={class:"gva-search-box"},X={class:"gva-table-box"},G={class:"gva-btn-list"},M={class:"popover-box"},W={key:1},Y={class:"popover-box"},Z={key:1},ee={class:"gva-pagination"},te=Object.assign({name:"SysOperationRecord"},{__name:"sysOperationRecord",setup(e){const t=r(1),a=r(0),l=r(10),o=r([]),i=r({}),c=()=>{i.value={}},v=()=>{t.value=1,l.value=10,""===i.value.status&&(i.value.status=null),h()},m=e=>{l.value=e,h()},g=e=>{t.value=e,h()},h=async()=>{const e=await(r={page:t.value,pageSize:l.value,...i.value},w({url:"/sysOperationRecord/getSysOperationRecordList",method:"get",params:r}));var r;0===e.code&&(o.value=e.data.list,a.value=e.data.total,t.value=e.data.page,l.value=e.data.pageSize)};h();const B=r([]),y=e=>{B.value=e},z=async()=>{D.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{const e=[];B.value&&B.value.forEach((t=>{e.push(t.ID)}));var a;0===(await(a={ids:e},w({url:"/sysOperationRecord/deleteSysOperationRecordByIds",method:"delete",data:a}))).code&&(j({type:"success",message:"删除成功"}),o.value.length===e.length&&t.value>1&&t.value--,h())}))},T=async e=>{D.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{var a;0===(await(a={ID:e.ID},w({url:"/sysOperationRecord/deleteSysOperationRecord",method:"delete",data:a}))).code&&(j({type:"success",message:"删除成功"}),1===o.value.length&&t.value>1&&t.value--,h())}))},F=e=>{try{return JSON.parse(e)}catch(t){return e}};return(e,r)=>{const h=x,w=V,D=k,j=I,$=R,P=U,q=C("warning"),H=S,J=L,K=N,te=O;return n(),u("div",null,[A("div",Q,[E(j,{inline:!0,model:i.value},{default:d((()=>[E(w,{label:"请求方法"},{default:d((()=>[E(h,{modelValue:i.value.method,"onUpdate:modelValue":r[0]||(r[0]=e=>i.value.method=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),E(w,{label:"请求路径"},{default:d((()=>[E(h,{modelValue:i.value.path,"onUpdate:modelValue":r[1]||(r[1]=e=>i.value.path=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),E(w,{label:"结果状态码"},{default:d((()=>[E(h,{modelValue:i.value.status,"onUpdate:modelValue":r[2]||(r[2]=e=>i.value.status=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),E(w,null,{default:d((()=>[E(D,{type:"primary",icon:"search",onClick:v},{default:d((()=>[b("查询")])),_:1}),E(D,{icon:"refresh",onClick:c},{default:d((()=>[b("重置")])),_:1})])),_:1})])),_:1},8,["model"])]),A("div",X,[A("div",G,[E(D,{icon:"delete",disabled:!B.value.length,onClick:z},{default:d((()=>[b("删除")])),_:1},8,["disabled"])]),E(K,{ref:"multipleTable",data:o.value,style:{width:"100%"},"tooltip-effect":"dark","row-key":"ID",onSelectionChange:y},{default:d((()=>[E($,{align:"left",type:"selection",width:"55"}),E($,{align:"left",label:"操作人",width:"140"},{default:d((e=>[A("div",null,f(e.row.user.userName)+"("+f(e.row.user.nickName)+")",1)])),_:1}),E($,{align:"left",label:"日期",width:"180"},{default:d((e=>[b(f(s(_)(e.row.CreatedAt)),1)])),_:1}),E($,{align:"left",label:"状态码",prop:"status",width:"120"},{default:d((e=>[A("div",null,[E(P,{type:"success"},{default:d((()=>[b(f(e.row.status),1)])),_:2},1024)])])),_:1}),E($,{align:"left",label:"请求IP",prop:"ip",width:"120"}),E($,{align:"left",label:"请求方法",prop:"method",width:"120"}),E($,{align:"left",label:"请求路径",prop:"path",width:"240"}),E($,{align:"left",label:"请求",prop:"path",width:"80"},{default:d((e=>[A("div",null,[e.row.body?(n(),p(J,{key:0,placement:"left-start"},{reference:d((()=>[E(H,{style:{cursor:"pointer"}},{default:d((()=>[E(q)])),_:1})])),default:d((()=>[A("div",M,[A("pre",null,f(F(e.row.body)),1)])])),_:2},1024)):(n(),u("span",W,"无"))])])),_:1}),E($,{align:"left",label:"响应",prop:"path",width:"80"},{default:d((e=>[A("div",null,[e.row.resp?(n(),p(J,{key:0,placement:"left-start"},{reference:d((()=>[E(H,{style:{cursor:"pointer"}},{default:d((()=>[E(q)])),_:1})])),default:d((()=>[A("div",Y,[A("pre",null,f(F(e.row.resp)),1)])])),_:2},1024)):(n(),u("span",Z,"无"))])])),_:1}),E($,{align:"left",label:"操作"},{default:d((e=>[E(D,{icon:"delete",type:"primary",link:"",onClick:t=>T(e.row)},{default:d((()=>[b("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),A("div",ee,[E(te,{"current-page":t.value,"page-size":l.value,"page-sizes":[10,30,50,100],total:a.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:g,onSizeChange:m},null,8,["current-page","page-size","total"])])])])}}});export{te as default};
