/*! 
 Build based on BAIDU 
 Time : 1716207707000 */
import{s as a,a as e,V as t,o as s,c as o,b as r,d as i,w as l,f as n,j as p,k as c,E as d}from"./087AC4D233B64EB0index.5fcd1422.js";import{E as u}from"./087AC4D233B64EB0scrollbar.b45aa0d9.js";import{E as h}from"./087AC4D233B64EB0tree.0c7e437e.js";import"./087AC4D233B64EB0checkbox.2864820c.js";/* empty css                               *//* empty css                              */import{e as m}from"./087AC4D233B64EB0api.a9a90911.js";import"./087AC4D233B64EB0index.fc2f110f.js";import"./087AC4D233B64EB0isEqual.c192b418.js";import"./087AC4D233B64EB0_Uint8Array.72188c2c.js";const f={class:"sticky top-0.5 z-10 bg-white"},y={class:"tree-content"},B=Object.assign({name:"Apis"},{__name:"apis",props:{row:{default:function(){return{}},type:Object}},setup(B,{expose:v}){const E=B,b=e({children:"children",label:"description"}),j=e(""),C=e([]),A=e([]),k=e("");(async()=>{const e=(await m()).data.apis;C.value=g(e);const t=await(s={authorityId:E.row.authorityId},a({url:"/casbin/getPolicyPathByAuthorityId",method:"post",data:s}));var s;k.value=E.row.authorityId,A.value=[],t.data.paths&&t.data.paths.forEach((a=>{A.value.push("p:"+a.path+"m:"+a.method)}))})();const w=e(!1),D=()=>{w.value=!0},g=a=>{const e={};a&&a.forEach((a=>{a.onlyId="p:"+a.path+"m:"+a.method,Object.prototype.hasOwnProperty.call(e,a.apiGroup)?e[a.apiGroup].push(a):Object.assign(e,{[a.apiGroup]:[a]})}));const t=[];for(const s in e){const a={ID:s,description:s+"组",children:e[s]};t.push(a)}return t},x=e(null),I=async()=>{const e=x.value.getCheckedNodes(!0);var t=[];e&&e.forEach((a=>{var e={path:a.path,method:a.method};t.push(e)}));var s;0===(await(s={authorityId:k.value,casbinInfos:t},a({url:"/casbin/updateCasbin",method:"post",data:s}))).code&&p({type:"success",message:"api设置成功"})};v({needConfirm:w,enterAndNext:()=>{I()}});const O=(a,e)=>!a||-1!==e.description.indexOf(a);return t(j,(a=>{x.value.filter(a)})),(a,e)=>{const t=c,p=d,m=h,B=u;return s(),o("div",null,[r("div",f,[i(t,{modelValue:j.value,"onUpdate:modelValue":e[0]||(e[0]=a=>j.value=a),class:"w-3/5",placeholder:"筛选"},null,8,["modelValue"]),i(p,{class:"float-right",type:"primary",onClick:I},{default:l((()=>[n("确 定")])),_:1})]),r("div",y,[i(B,null,{default:l((()=>[i(m,{ref_key:"apiTree",ref:x,data:C.value,"default-checked-keys":A.value,props:b.value,"default-expand-all":"","highlight-current":"","node-key":"onlyId","show-checkbox":"","filter-node-method":O,onCheck:D},null,8,["data","default-checked-keys","props"])])),_:1})])])}}});export{B as default};
