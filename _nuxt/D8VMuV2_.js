import{d as S,ao as A,ap as f,b as o,c as p,n as u,E as t,ai as x,j as L,t as $,f as m,e as b,g as C,F as T,ag as R,ah as v,aj as _,ab as j,an as le,w as O,aE as ce,l as ue,al as de,am as Z,ac as V,D as U,aF as M,as as E,y as D,aG as pe,aH as ge,r as H,H as me,aI as ye,aJ as be,u as fe,aK as ve,k as he,aL as ke,v as xe,aM as ee,aN as _e,aO as we,av as te,aP as ae,aQ as $e,s as ne,aR as Ce}from"./B_z3COam.js";import Se from"./CIZkk4Fk.js";import{a as Ae}from"./BhCHo9Hz.js";import"./DasOLl54.js";import"./D_HsUiuO.js";import"./C-F1gjmX.js";const Ue={wrapper:{base:"flex items-center align-center text-center",horizontal:"w-full flex-row",vertical:"flex-col"},container:{base:"font-medium text-gray-700 dark:text-gray-200 flex",horizontal:"mx-3 whitespace-nowrap",vertical:"my-2"},border:{base:"flex border-gray-200 dark:border-gray-800",horizontal:"w-full",vertical:"h-full",size:{horizontal:{"2xs":"border-t",xs:"border-t-[2px]",sm:"border-t-[3px]",md:"border-t-[4px]",lg:"border-t-[5px]",xl:"border-t-[6px]"},vertical:{"2xs":"border-s",xs:"border-s-[2px]",sm:"border-s-[3px]",md:"border-s-[4px]",lg:"border-s-[5px]",xl:"border-s-[6px]"}},type:{solid:"border-solid",dotted:"border-dotted",dashed:"border-dashed"}},icon:{base:"flex-shrink-0 w-5 h-5"},avatar:{base:"flex-shrink-0",size:"2xs"},label:"text-sm",default:{size:"2xs"}},Oe={class:"flex flex-col lg:flex-row items-start gap-6"},je={class:"flex-1"},ze=S({inheritAttrs:!1,__name:"PageHeader",props:{headline:{type:String,default:void 0},title:{type:String,default:void 0},description:{type:String,default:void 0},icon:{type:String,default:void 0},links:{type:Array,default:()=>[]},class:{type:[String,Object,Array],default:void 0},ui:{type:Object,default:()=>({})}},setup(e){const s={wrapper:"relative border-b border-gray-200 dark:border-gray-800 py-8",container:"flex flex-col lg:flex-row lg:items-center lg:justify-between",headline:"mb-3 text-sm/6 font-semibold text-primary flex items-center gap-1.5",title:"text-3xl sm:text-4xl font-bold text-gray-900 dark:text-white tracking-tight",description:"mt-4 text-lg text-gray-500 dark:text-gray-400",icon:{wrapper:"flex",base:"w-10 h-10 flex-shrink-0 text-primary"},links:"flex flex-wrap items-center gap-1.5 mt-4 lg:mt-0"},l=e,{ui:r,attrs:a}=A("page.header",f(l,"ui"),s,f(l,"class"),!0);return(n,c)=>{var g;const i=j,y=le;return o(),p("div",_({class:t(r).wrapper},t(a)),[e.headline||n.$slots.headline?(o(),p("div",{key:0,class:u(t(r).headline)},[x(n.$slots,"headline",{},()=>[L($(e.headline),1)])],2)):m("",!0),b("div",Oe,[e.icon||n.$slots.icon?(o(),p("div",{key:0,class:u(t(r).icon.wrapper)},[x(n.$slots,"icon",{},()=>[C(i,{name:e.icon,class:u(t(r).icon.base)},null,8,["name","class"])])],2)):m("",!0),b("div",je,[b("div",{class:u(t(r).container)},[b("h1",{class:u(t(r).title)},[x(n.$slots,"title",{},()=>[L($(e.title),1)])],2),(g=e.links)!=null&&g.length||n.$slots.links?(o(),p("div",{key:0,class:u(t(r).links)},[x(n.$slots,"links",{},()=>[(o(!0),p(T,null,R(e.links,(h,d)=>(o(),v(y,_({key:d},{...h,target:h.target||"_blank",color:h.color||"white"},{onClick:h.click}),null,16,["onClick"]))),128))])],2)):m("",!0)],2),e.description||n.$slots.description?(o(),p("p",{key:0,class:u(t(r).description)},[x(n.$slots,"description",{},()=>[L($(e.description),1)])],2)):m("",!0),x(n.$slots,"default")])])],16)}}}),Ie=S({inheritAttrs:!1,__name:"ContentSurroundLink",props:{link:{type:Object,required:!0},icon:{type:String,default:void 0},class:{type:[String,Object,Array],default:void 0},ui:{type:Object,default:()=>({})}},setup(e){const s={wrapper:"block px-6 py-8 border not-prose rounded-lg border-gray-200 dark:border-gray-800 hover:bg-gray-100/50 dark:hover:bg-gray-800/50 group",icon:{wrapper:"inline-flex items-center rounded-full p-1.5 bg-gray-100 dark:bg-gray-800 group-hover:bg-primary/10 ring-1 ring-gray-300 dark:ring-gray-700 mb-4 group-hover:ring-primary/50",base:"w-5 h-5 text-gray-900 dark:text-white group-hover:text-primary"},title:"font-medium text-gray-900 dark:text-white text-[15px] mb-1",description:"text-sm font-normal text-gray-500 dark:text-gray-400 line-clamp-2"},l=e,{ui:r,attrs:a}=A("content.surround.link",f(l,"ui"),s,f(l,"class"),!0);return(n,c)=>{const i=j,y=ce;return o(),v(y,_({to:e.link._path,class:t(r).wrapper},t(a)),{default:O(()=>[e.icon||e.link.icon?(o(),p("div",{key:0,class:u(t(r).icon.wrapper)},[C(i,{name:e.icon||e.link.icon,class:u(t(r).icon.base)},null,8,["name","class"])],2)):m("",!0),b("p",{class:u(t(r).title)},$(e.link.title),3),b("p",{class:u(t(r).description)},$(e.link.description),3)]),_:1},16,["to","class"])}}}),Le={key:1,class:"hidden sm:block"},Pe=S({inheritAttrs:!1,__name:"ContentSurround",props:{surround:{type:Array,default:()=>[]},class:{type:[String,Object,Array],default:void 0},ui:{type:Object,default:()=>({})}},setup(e){const s={wrapper:"grid gap-8 sm:grid-cols-2",icon:{prev:"i-heroicons-arrow-left-20-solid",next:"i-heroicons-arrow-right-20-solid"},link:{}},l=e,{ui:r,attrs:a}=A("content.surround",f(l,"ui"),s,f(l,"class"),!0),[n,c]=l.surround||[];return(i,y)=>{const g=Ie;return o(),p("div",_({class:t(r).wrapper},t(a)),[t(n)?(o(),v(g,{key:0,link:t(n),ui:t(r).link,icon:t(r).icon.prev},null,8,["link","ui","icon"])):(o(),p("span",Le," ")),t(c)?(o(),v(g,{key:2,link:t(c),ui:t(r).link,icon:t(r).icon.next,class:"text-right"},null,8,["link","ui","icon"])):m("",!0)],16)}}}),Be=S({inheritAttrs:!1,__name:"PageBody",props:{prose:{type:Boolean,default:!1},class:{type:[String,Object,Array],default:void 0},ui:{type:Object,default:()=>({})}},setup(e){const s={wrapper:"mt-8 pb-24",prose:"prose prose-primary dark:prose-invert max-w-none"},l=e,{ui:r,attrs:a}=A("page.body",f(l,"ui"),s,f(l,"class"),!0);return(n,c)=>(o(),p("div",_({class:[t(r).wrapper,e.prose&&t(r).prose]},t(a)),[x(n.$slots,"default")],16))}}),B=de(Z.ui.strategy,Z.ui.divider,Ue),He=S({components:{UIcon:j,UAvatar:V},inheritAttrs:!1,props:{label:{type:String,default:null},icon:{type:String,default:null},avatar:{type:Object,default:null},size:{type:String,default:()=>B.default.size,validator(e){return Object.keys(B.border.size.horizontal).includes(e)||Object.keys(B.border.size.vertical).includes(e)}},orientation:{type:String,default:"horizontal",validator:e=>["horizontal","vertical"].includes(e)},type:{type:String,default:"solid",validator:e=>["solid","dotted","dashed"].includes(e)},class:{type:[String,Object,Array],default:()=>""},ui:{type:Object,default:()=>({})}},setup(e){const{ui:s,attrs:l}=A("divider",f(e,"ui"),B),r=U(()=>M(E(s.value.wrapper.base,s.value.wrapper[e.orientation]),e.class)),a=U(()=>E(s.value.container.base,s.value.container[e.orientation])),n=U(()=>E(s.value.border.base,s.value.border[e.orientation],s.value.border.size[e.orientation][e.size],s.value.border.type[e.type]));return{ui:s,attrs:l,wrapperClass:r,containerClass:a,borderClass:n}}});function Te(e,s,l,r,a,n){const c=j,i=V;return o(),p("div",_({class:e.wrapperClass},e.attrs),[b("div",{class:u(e.borderClass)},null,2),e.label||e.icon||e.avatar||e.$slots.default?(o(),p(T,{key:0},[b("div",{class:u(e.containerClass)},[x(e.$slots,"default",{},()=>[e.label?(o(),p("span",{key:0,class:u(e.ui.label)},$(e.label),3)):e.icon?(o(),v(c,{key:1,name:e.icon,class:u(e.ui.icon.base)},null,8,["name","class"])):e.avatar?(o(),v(i,_({key:2},{size:e.ui.avatar.size,...e.avatar},{class:e.ui.avatar.base}),null,16,["class"])):m("",!0)])],2),b("div",{class:u(e.borderClass)},null,2)],64)):m("",!0)],16)}const Ne=ue(He,[["render",Te]]),Ee=S({inheritAttrs:!1,__name:"PageLinks",props:{title:{type:String,default:void 0},links:{type:Array,default:()=>[]},class:{type:[String,Object,Array],default:void 0},ui:{type:Object,default:()=>({})}},setup(e){const s=D(),l=U(()=>({wrapper:"space-y-3",title:"text-sm/6 font-semibold flex items-center gap-1.5",container:"space-y-3 lg:space-y-1.5",base:"flex items-center gap-1.5",active:"text-primary",inactive:"text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200",icon:{base:"w-5 h-5 flex-shrink-0"},avatar:{base:"self-center",size:"2xs"},externalIcon:{name:s.ui.icons.external,base:"w-3 h-3 absolute top-0.5 -right-3.5 text-gray-400 dark:text-gray-500"},label:"text-sm/6 font-medium relative"})),r=e,{ui:a,attrs:n}=A("page.links",f(r,"ui"),l,f(r,"class"),!0);return(c,i)=>{const y=j,g=V,h=ge;return o(),p("div",_({class:t(a).wrapper},t(n)),[e.title||c.$slots.title?(o(),p("p",{key:0,class:u(t(a).title)},[x(c.$slots,"title",{},()=>[L($(e.title),1)])],2)):m("",!0),b("div",{class:u(t(a).container)},[x(c.$slots,"default",{},()=>[(o(!0),p(T,null,R(e.links,(d,k)=>(o(),v(h,_({key:k},t(pe)(d),{class:t(a).base,"active-class":t(a).active,"inactive-class":t(a).inactive,onClick:d.click}),{default:O(()=>[d.icon?(o(),v(y,{key:0,name:d.icon,class:u(t(M)(t(a).icon.base,d.iconClass))},null,8,["name","class"])):m("",!0),d.avatar?(o(),v(g,_({key:1},{size:t(a).avatar.size,...d.avatar},{class:t(M)(t(a).avatar.base,d.avatarClass)}),null,16,["class"])):m("",!0),b("span",{class:u(t(a).label)},[L($(d.label)+" ",1),d.target==="_blank"?(o(),v(y,{key:0,name:t(a).externalIcon.name,class:u(t(a).externalIcon.base)},null,8,["name","class"])):m("",!0)],2)]),_:2},1040,["class","active-class","inactive-class","onClick"]))),128))])],2)],16)}}}),Me=()=>{const e=H(),s=H([]),l=H([]),r=n=>{n.forEach(c=>{const i=c.target.id;c.isIntersecting?s.value=[...s.value,i]:s.value=s.value.filter(y=>y!==i)})},a=n=>{n.forEach(c=>{e.value&&e.value.observe(c)})};return me(s,(n,c)=>{n.length===0?l.value=c:l.value=n}),ye(()=>e.value=new IntersectionObserver(r)),be(()=>{var n;return(n=e.value)==null?void 0:n.disconnect()}),{visibleHeadings:s,activeHeadings:l,updateHeadings:a}},Re=["href","onClick"],se=S({inheritAttrs:!1,__name:"ContentTocLinks",props:{links:{type:Array,default:()=>[]},class:{type:[String,Object,Array],default:void 0},ui:{type:Object,default:()=>({})}},emits:["move"],setup(e,{emit:s}){const l={wrapper:"space-y-1",base:"block text-sm/6 truncate",active:"text-primary",inactive:"text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200",depth:"ml-3"},r=e,a=s,n=fe(),c=he(),{activeHeadings:i,updateHeadings:y}=Me(),{ui:g,attrs:h}=A("content.toc.links",f(r,"ui"),l,f(r,"class"),!0);c.hooks.hookOnce("page:finish",()=>{y([...document.querySelectorAll("h2"),...document.querySelectorAll("h3")])});const d=k=>{const z=encodeURIComponent(k);n.push(`#${z}`),a("move",k)};return(k,z)=>{var P;const N=se;return(P=e.links)!=null&&P.length?(o(),p("ul",_({key:0,class:t(g).wrapper},t(h)),[(o(!0),p(T,null,R(e.links,w=>(o(),p("li",{key:w.text,class:u([t(g).wrapper,w.depth===3&&t(g).depth])},[b("a",{href:`#${w.id}`,class:u([t(g).base,t(i).includes(w.id)?t(g).active:t(g).inactive]),onClick:ve(q=>d(w.id),["prevent"])},$(w.text),11,Re),w.children?(o(),v(N,{key:0,links:w.children},null,8,["links"])):m("",!0)],2))),128))],16)):m("",!0)}}}),Ve=S({inheritAttrs:!1,__name:"ContentToc",props:{title:{type:String,default:"Table of Contents"},links:{type:Array,default:()=>[]},class:{type:[String,Object,Array],default:void 0},ui:{type:Object,default:()=>({})}},setup(e){const s=D(),l=U(()=>({wrapper:"sticky top-[--header-height] bg-background/75 backdrop-blur -mx-4 sm:-mx-6 px-4 sm:px-6 lg:px-4 lg:-mx-4 overflow-y-auto max-h-[calc(100vh-var(--header-height))]",container:{base:"py-3 lg:py-8 border-b border-dashed border-gray-200 dark:border-gray-800 lg:border-0 space-y-3",empty:"lg:py-8 space-y-3"},button:{base:"flex items-center gap-1.5 lg:cursor-text lg:select-text w-full group",label:"font-semibold text-sm/6 truncate",trailingIcon:{name:s.ui.icons.chevron,base:"w-5 h-5 ms-auto transform transition-transform duration-200 flex-shrink-0 mr-1.5",active:"text-gray-700 dark:text-gray-200",inactive:"text-gray-500 dark:text-gray-400 group-hover:text-gray-700 dark:group-hover:text-gray-200 -rotate-90"}},links:{}})),r=e,{ui:a,attrs:n}=A("content.toc",f(r,"ui"),l,f(r,"class"),!0),c=H(!1);return(i,y)=>{var d,k;const g=j,h=se;return o(),p("nav",_({class:t(a).wrapper},t(n)),[b("div",{class:u([(d=e.links)!=null&&d.length?t(a).container.base:t(a).container.empty])},[x(i.$slots,"top"),(k=e.links)!=null&&k.length?(o(),p("button",{key:0,class:u(t(a).button.base),tabindex:"-1",onClick:y[0]||(y[0]=z=>c.value=!t(c))},[b("span",{class:u(t(a).button.label)},$(e.title),3),C(g,{name:t(a).button.trailingIcon.name,class:u(["lg:!hidden",[t(a).button.trailingIcon.base,t(c)?t(a).button.trailingIcon.active:t(a).button.trailingIcon.inactive]])},null,8,["name","class"])],2)):m("",!0),C(h,{links:e.links,ui:t(a).links,class:u([t(c)?"lg:block":"hidden lg:block"])},null,8,["links","ui","class"]),x(i.$slots,"bottom")],2)],16)}}}),De={key:1},We=S({__name:"[...slug]",async setup(e){let s,l;const{t:r}=ke(),a=xe(),{toc:n,seo:c}=D(),{data:i}=([s,l]=ee(()=>ae(a.path,()=>ne(a.path).findOne(),"$tVA6J66VEN")),s=await s,l(),s);if(!i.value)throw _e({statusCode:404,statusMessage:"Page not found",fatal:!0});const{data:y}=([s,l]=ee(()=>ae(`${a.path}-surround`,()=>ne().where({_extension:"md",navigation:{$ne:!1}}).only(["title","description","_path"]).findSurround(Ce(a.path)))),s=await s,l(),s);we({title:i.value.title,ogTitle:`${i.value.title} - ${c==null?void 0:c.siteName}`,description:i.value.description,ogDescription:i.value.description});const g=U(()=>Ae(i.value)),h=U(()=>{var d,k;return[((d=n==null?void 0:n.bottom)==null?void 0:d.edit)&&{icon:"i-mdi-pencil-box",label:r("docs.toc.bottom.edit"),to:`${n.bottom.edit}/${(k=i==null?void 0:i.value)==null?void 0:k._file}`,target:"_blank"},{icon:"i-mdi-star",label:r("docs.toc.bottom.star"),to:"https://github.com/nuxt/ui",target:"_blank"}].filter(Boolean)});return(d,k)=>{const z=ze,N=Se,P=Pe,w=Be,q=Ne,re=Ee,oe=Ve,ie=$e;return o(),v(ie,null,te({default:O(()=>[C(z,{title:t(i).title,description:t(i).description,links:t(i).links,headline:t(g)},null,8,["title","description","links","headline"]),C(w,{prose:""},{default:O(()=>{var I;return[t(i).body?(o(),v(N,{key:0,value:t(i)},null,8,["value"])):m("",!0),(I=t(y))!=null&&I.length?(o(),p("hr",De)):m("",!0),C(P,{surround:t(y)},null,8,["surround"])]}),_:1})]),_:2},[t(i).toc!==!1?{name:"right",fn:O(()=>{var I,F,J;return[C(oe,{title:d.$t("common.toc"),links:(F=(I=t(i).body)==null?void 0:I.toc)==null?void 0:F.links},te({_:2},[(J=t(n))!=null&&J.bottom?{name:"bottom",fn:O(()=>{var G,K,Q,W,X,Y;return[b("div",{class:u(["hidden space-y-6 lg:block",{"!mt-6":(Q=(K=(G=t(i).body)==null?void 0:G.toc)==null?void 0:K.links)==null?void 0:Q.length}])},[(Y=(X=(W=t(i).body)==null?void 0:W.toc)==null?void 0:X.links)!=null&&Y.length?(o(),v(q,{key:0,type:"dashed"})):m("",!0),C(re,{title:d.$t("common.community"),links:t(h)},null,8,["title","links"])],2)]}),key:"0"}:void 0]),1032,["title","links"])]}),key:"0"}:void 0]),1024)}}});export{We as default};
