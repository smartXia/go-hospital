/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{B as e,C as o,ad as l,bm as a,b4 as n,bG as t,aW as s,a as u,cc as c,cd as r,D as d,aX as i,ce as p,V as f,ak as y,X as v,a2 as m,aM as B,bk as C}from"./087AC4D233B64EB0index.5fcd1422.js";import{i as b}from"./087AC4D233B64EB0index.71b3c6ca.js";const g=e({center:Boolean,alignCenter:Boolean,closeIcon:{type:o},draggable:Boolean,overflow:Boolean,fullscreen:Boolean,showClose:{type:Boolean,default:!0},title:{type:String,default:""},ariaLevel:{type:String,default:"2"}}),D={close:()=>!0},x=e({...g,appendToBody:Boolean,appendTo:{type:l(String),default:"body"},beforeClose:{type:l(Function)},destroyOnClose:Boolean,closeOnClickModal:{type:Boolean,default:!0},closeOnPressEscape:{type:Boolean,default:!0},lockScroll:{type:Boolean,default:!0},modal:{type:Boolean,default:!0},openDelay:{type:Number,default:0},closeDelay:{type:Number,default:0},top:{type:String},modelValue:Boolean,modalClass:String,width:{type:[String,Number]},zIndex:{type:Number},trapFocus:{type:Boolean,default:!1},headerAriaLevel:{type:String,default:"2"}}),S={open:()=>!0,opened:()=>!0,close:()=>!0,closed:()=>!0,[a]:e=>n(e),openAutoFocus:()=>!0,closeAutoFocus:()=>!0},A=(e,o)=>{var l;const n=m().emit,{nextZIndex:g}=t();let D="";const x=s(),S=s(),A=u(!1),F=u(!1),I=u(!1),h=u(null!=(l=e.zIndex)?l:g());let k,O;const w=c("namespace",r),E=d((()=>{const o={},l="--".concat(w.value,"-dialog");return e.fullscreen||(e.top&&(o["".concat(l,"-margin-top")]=e.top),e.width&&(o["".concat(l,"-width")]=i(e.width))),o})),z=d((()=>e.alignCenter?{display:"flex"}:{}));function L(){null==O||O(),null==k||k(),e.openDelay&&e.openDelay>0?({stop:k}=B((()=>V()),e.openDelay)):V()}function M(){null==k||k(),null==O||O(),e.closeDelay&&e.closeDelay>0?({stop:O}=B((()=>P()),e.closeDelay)):P()}function N(){e.beforeClose?e.beforeClose((function(e){e||(F.value=!0,A.value=!1)})):M()}function V(){C&&(A.value=!0)}function P(){A.value=!1}return e.lockScroll&&p(A),f((()=>e.modelValue),(l=>{l?(F.value=!1,L(),I.value=!0,h.value=b(e.zIndex)?g():h.value++,y((()=>{n("open"),o.value&&(o.value.scrollTop=0)}))):A.value&&M()})),f((()=>e.fullscreen),(e=>{o.value&&(e?(D=o.value.style.transform,o.value.style.transform=""):o.value.style.transform=D)})),v((()=>{e.modelValue&&(A.value=!0,I.value=!0,L())})),{afterEnter:function(){n("opened")},afterLeave:function(){n("closed"),n(a,!1),e.destroyOnClose&&(I.value=!1)},beforeLeave:function(){n("close")},handleClose:N,onModalClick:function(){e.closeOnClickModal&&N()},close:M,doClose:P,onOpenAutoFocus:function(){n("openAutoFocus")},onCloseAutoFocus:function(){n("closeAutoFocus")},onCloseRequested:function(){e.closeOnPressEscape&&N()},onFocusoutPrevented:function(e){var o;"pointer"===(null==(o=e.detail)?void 0:o.focusReason)&&e.preventDefault()},titleId:x,bodyId:S,closed:F,style:E,overlayDialogStyle:z,rendered:I,visible:A,zIndex:h}};export{S as a,g as b,D as c,x as d,A as u};
