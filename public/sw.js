if(!self.define){let e,a={};const i=(i,t)=>(i=new URL(i+".js",t).href,a[i]||new Promise((a=>{if("document"in self){const e=document.createElement("script");e.src=i,e.onload=a,document.head.appendChild(e)}else e=i,importScripts(i),a()})).then((()=>{let e=a[i];if(!e)throw new Error(`Module ${i} didn’t register its module`);return e})));self.define=(t,s)=>{const n=e||("document"in self?document.currentScript.src:"")||location.href;if(a[n])return;let c={};const f=e=>i(e,n),r={module:{uri:n},exports:c,require:f};a[n]=Promise.all(t.map((e=>r[e]||f(e)))).then((e=>(s(...e),c)))}}define(["./workbox-50de5c5d"],(function(e){"use strict";importScripts(),self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"/_next/app-build-manifest.json",revision:"27bf6510efda9a8d21856edfaebcbf13"},{url:"/_next/static/bAIw2Q76HnJoDf3T1MNT8/_buildManifest.js",revision:"75740cacd3ef418c900cdf5afc2f6581"},{url:"/_next/static/bAIw2Q76HnJoDf3T1MNT8/_ssgManifest.js",revision:"b6652df95db52feb4daf4eca35380933"},{url:"/_next/static/chunks/00cbbcb7-f6acf453502df7c6.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/08ffe114-9b7cae2ac07d7712.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/1124-26e1cd342880efe5.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/1179-ba7408a497468a04.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/1259.a76ae9932b86abf6.js",revision:"a76ae9932b86abf6"},{url:"/_next/static/chunks/2446-5a9d91b46fd1e8f0.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/3360-82cb5bdb4005b320.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/3627521c-851f851e7494f58d.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/4292-3ec7520ff3af714e.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/4445-792d25077a50272b.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/4724-e69d9aeef37e0aac.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/4743-243c0409b9049a5e.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/4997-9f32a00778a46eb7.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/5365-e2264c8ae389dc88.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/5764-fecf35c646252401.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/6107-aa90381f673a2c72.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/6398-5103e39580899143.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/660-872ddc309987a454.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/6691-fe46492252832676.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/6708-a2af16f10fe7bf46.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/7021-c853ac5fd9886550.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/7826-aec7ee5d12b46612.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/8740-d22ff4b49711ef8e.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/8dc5345f-496255aaffd2df18.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/9006-48d3fcc2f5428294.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/9081a741-f2f2de4799081c49.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/9304-5a226e6a8d45197d.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/9363-5d7d8a8beae80493.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/9530-f5174b66f330960a.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/9806-6013e11161529768.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/9866-9138f6b898fe1d08.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/auth/layout-15a2c7dec9aa9306.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/auth/sign-in/%5B%5B...sign-in%5D%5D/page-48e115cc95f98a57.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/auth/sign-up/%5B%5B...sign-up%5D%5D/page-2c671d117d749519.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/blog/%5Bslug%5D/page-21ea85c4a85a7f8a.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/blog/layout-dfa13adf7906bdd6.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/blog/page-61b56be3075ae010.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/certificate/layout-04cba613c0f90a88.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/certificate/page-9875b198fc4e346f.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/error-1eea53003527fe86.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/ia/chat-pdf/%5Bid%5D/page-7b08c551adf7cbfb.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/ia/chat-pdf/layout-86f3b5da27e00e8e.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/ia/chat-pdf/page-e7153cce1f7a7161.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/ia/layout-f0fd0c0e2e8a458f.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/ia/page-400ab123ef530d18.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/layout-d92f64640613aca5.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/loading-511a91860eae9c1c.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/me/page-90f741b1bb6e1a0e.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/not-found-c36bbc1851a80f3f.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/page-85153b3a411151d3.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/project/%5Bslug%5D/page-780bb1e53701b356.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/project/layout-b7b11d7290ada38f.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/app/project/page-b34a63b124529be5.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/bc9c3264-cc97c0e0c3e60060.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/bf6a786c-98e8ed74990cff98.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/framework-4498e84bb0ba1830.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/main-app-55c28f89475dcad4.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/main-fc51ba4944fa8d58.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/pages/_app-7bb460e314c5f602.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/pages/_error-8aa332dfaf8ff0ba.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js",revision:"837c0df77fd5009c9e46d446188ecfd0"},{url:"/_next/static/chunks/webpack-e8d6ae81b50992c5.js",revision:"bAIw2Q76HnJoDf3T1MNT8"},{url:"/_next/static/css/2aa199ae6265e4c0.css",revision:"2aa199ae6265e4c0"},{url:"/_next/static/css/3ffdcef73c138c9b.css",revision:"3ffdcef73c138c9b"},{url:"/_next/static/css/b9b07d7c9a83825f.css",revision:"b9b07d7c9a83825f"},{url:"/_next/static/css/cf7b59adfe932c8d.css",revision:"cf7b59adfe932c8d"},{url:"/_next/static/media/0e4fe491bf84089c-s.p.woff2",revision:"5e22a46c04d947a36ea0cad07afcc9e1"},{url:"/_next/static/media/1c57ca6f5208a29b-s.woff2",revision:"491a7a9678c3cfd4f86c092c68480f23"},{url:"/_next/static/media/37b0c0a51409261e-s.woff2",revision:"5ce748f413aee42a8d4723df0d18830b"},{url:"/_next/static/media/3dbd163d3bb09d47-s.woff2",revision:"93dcb0c222437699e9dd591d8b5a6b85"},{url:"/_next/static/media/42d52f46a26971a3-s.woff2",revision:"b44d0dd122f9146504d444f290252d88"},{url:"/_next/static/media/44c3f6d12248be7f-s.woff2",revision:"705e5297b1a92dac3b13b2705b7156a7"},{url:"/_next/static/media/46c894be853ec49f-s.woff2",revision:"47891b6adb3a947dd3c594bd5196850e"},{url:"/_next/static/media/4a8324e71b197806-s.woff2",revision:"5fba57b10417c946c556545c9f348bbd"},{url:"/_next/static/media/506bd11311670951-s.woff2",revision:"7976a92314c8770252603e7813da9f67"},{url:"/_next/static/media/5647e4c23315a2d2-s.woff2",revision:"e64969a373d0acf2586d1fd4224abb90"},{url:"/_next/static/media/627622453ef56b0d-s.p.woff2",revision:"e7df3d0942815909add8f9d0c40d00d9"},{url:"/_next/static/media/71ba03c5176fbd9c-s.woff2",revision:"2effa1fe2d0dff3e7b8c35ee120e0d05"},{url:"/_next/static/media/7be645d133f3ee22-s.woff2",revision:"3ba6fb27a0ea92c2f1513add6dbddf37"},{url:"/_next/static/media/7c53f7419436e04b-s.woff2",revision:"fd4ff709e3581e3f62e40e90260a1ad7"},{url:"/_next/static/media/7d8c9b0ca4a64a5a-s.p.woff2",revision:"0772a436bbaaaf4381e9d87bab168217"},{url:"/_next/static/media/80a2a8cc25a3c264-s.woff2",revision:"2d3d8a78ef164ab6c1c62a3e57c2727b"},{url:"/_next/static/media/83e4d81063b4b659-s.woff2",revision:"bd30db6b297b76f3a3a76f8d8ec5aac9"},{url:"/_next/static/media/8db47a8bf03b7d2f-s.p.woff2",revision:"49003e0ff09f1efb8323cf35b836ba8f"},{url:"/_next/static/media/8fb72f69fba4e3d2-s.woff2",revision:"7a2e2eae214e49b4333030f789100720"},{url:"/_next/static/media/912a9cfe43c928d9-s.woff2",revision:"376ffe2ca0b038d08d5e582ec13a310f"},{url:"/_next/static/media/934c4b7cb736f2a3-s.p.woff2",revision:"1f6d3cf6d38f25d83d95f5a800b8cac3"},{url:"/_next/static/media/94300924a0693016-s.woff2",revision:"105927314bd3f089b99c0dda456171ed"},{url:"/_next/static/media/9e48537b1b020091-s.woff2",revision:"4b52fd954ca934c204d73ddbc640e5d4"},{url:"/_next/static/media/KaTeX_AMS-Regular.1608a09b.woff",revision:"1608a09b"},{url:"/_next/static/media/KaTeX_AMS-Regular.4aafdb68.ttf",revision:"4aafdb68"},{url:"/_next/static/media/KaTeX_AMS-Regular.a79f1c31.woff2",revision:"a79f1c31"},{url:"/_next/static/media/KaTeX_Caligraphic-Bold.b6770918.woff",revision:"b6770918"},{url:"/_next/static/media/KaTeX_Caligraphic-Bold.cce5b8ec.ttf",revision:"cce5b8ec"},{url:"/_next/static/media/KaTeX_Caligraphic-Bold.ec17d132.woff2",revision:"ec17d132"},{url:"/_next/static/media/KaTeX_Caligraphic-Regular.07ef19e7.ttf",revision:"07ef19e7"},{url:"/_next/static/media/KaTeX_Caligraphic-Regular.55fac258.woff2",revision:"55fac258"},{url:"/_next/static/media/KaTeX_Caligraphic-Regular.dad44a7f.woff",revision:"dad44a7f"},{url:"/_next/static/media/KaTeX_Fraktur-Bold.9f256b85.woff",revision:"9f256b85"},{url:"/_next/static/media/KaTeX_Fraktur-Bold.b18f59e1.ttf",revision:"b18f59e1"},{url:"/_next/static/media/KaTeX_Fraktur-Bold.d42a5579.woff2",revision:"d42a5579"},{url:"/_next/static/media/KaTeX_Fraktur-Regular.7c187121.woff",revision:"7c187121"},{url:"/_next/static/media/KaTeX_Fraktur-Regular.d3c882a6.woff2",revision:"d3c882a6"},{url:"/_next/static/media/KaTeX_Fraktur-Regular.ed38e79f.ttf",revision:"ed38e79f"},{url:"/_next/static/media/KaTeX_Main-Bold.b74a1a8b.ttf",revision:"b74a1a8b"},{url:"/_next/static/media/KaTeX_Main-Bold.c3fb5ac2.woff2",revision:"c3fb5ac2"},{url:"/_next/static/media/KaTeX_Main-Bold.d181c465.woff",revision:"d181c465"},{url:"/_next/static/media/KaTeX_Main-BoldItalic.6f2bb1df.woff2",revision:"6f2bb1df"},{url:"/_next/static/media/KaTeX_Main-BoldItalic.70d8b0a5.ttf",revision:"70d8b0a5"},{url:"/_next/static/media/KaTeX_Main-BoldItalic.e3f82f9d.woff",revision:"e3f82f9d"},{url:"/_next/static/media/KaTeX_Main-Italic.47373d1e.ttf",revision:"47373d1e"},{url:"/_next/static/media/KaTeX_Main-Italic.8916142b.woff2",revision:"8916142b"},{url:"/_next/static/media/KaTeX_Main-Italic.9024d815.woff",revision:"9024d815"},{url:"/_next/static/media/KaTeX_Main-Regular.0462f03b.woff2",revision:"0462f03b"},{url:"/_next/static/media/KaTeX_Main-Regular.7f51fe03.woff",revision:"7f51fe03"},{url:"/_next/static/media/KaTeX_Main-Regular.b7f8fe9b.ttf",revision:"b7f8fe9b"},{url:"/_next/static/media/KaTeX_Math-BoldItalic.572d331f.woff2",revision:"572d331f"},{url:"/_next/static/media/KaTeX_Math-BoldItalic.a879cf83.ttf",revision:"a879cf83"},{url:"/_next/static/media/KaTeX_Math-BoldItalic.f1035d8d.woff",revision:"f1035d8d"},{url:"/_next/static/media/KaTeX_Math-Italic.5295ba48.woff",revision:"5295ba48"},{url:"/_next/static/media/KaTeX_Math-Italic.939bc644.ttf",revision:"939bc644"},{url:"/_next/static/media/KaTeX_Math-Italic.f28c23ac.woff2",revision:"f28c23ac"},{url:"/_next/static/media/KaTeX_SansSerif-Bold.8c5b5494.woff2",revision:"8c5b5494"},{url:"/_next/static/media/KaTeX_SansSerif-Bold.94e1e8dc.ttf",revision:"94e1e8dc"},{url:"/_next/static/media/KaTeX_SansSerif-Bold.bf59d231.woff",revision:"bf59d231"},{url:"/_next/static/media/KaTeX_SansSerif-Italic.3b1e59b3.woff2",revision:"3b1e59b3"},{url:"/_next/static/media/KaTeX_SansSerif-Italic.7c9bc82b.woff",revision:"7c9bc82b"},{url:"/_next/static/media/KaTeX_SansSerif-Italic.b4c20c84.ttf",revision:"b4c20c84"},{url:"/_next/static/media/KaTeX_SansSerif-Regular.74048478.woff",revision:"74048478"},{url:"/_next/static/media/KaTeX_SansSerif-Regular.ba21ed5f.woff2",revision:"ba21ed5f"},{url:"/_next/static/media/KaTeX_SansSerif-Regular.d4d7ba48.ttf",revision:"d4d7ba48"},{url:"/_next/static/media/KaTeX_Script-Regular.03e9641d.woff2",revision:"03e9641d"},{url:"/_next/static/media/KaTeX_Script-Regular.07505710.woff",revision:"07505710"},{url:"/_next/static/media/KaTeX_Script-Regular.fe9cbbe1.ttf",revision:"fe9cbbe1"},{url:"/_next/static/media/KaTeX_Size1-Regular.e1e279cb.woff",revision:"e1e279cb"},{url:"/_next/static/media/KaTeX_Size1-Regular.eae34984.woff2",revision:"eae34984"},{url:"/_next/static/media/KaTeX_Size1-Regular.fabc004a.ttf",revision:"fabc004a"},{url:"/_next/static/media/KaTeX_Size2-Regular.57727022.woff",revision:"57727022"},{url:"/_next/static/media/KaTeX_Size2-Regular.5916a24f.woff2",revision:"5916a24f"},{url:"/_next/static/media/KaTeX_Size2-Regular.d6b476ec.ttf",revision:"d6b476ec"},{url:"/_next/static/media/KaTeX_Size3-Regular.9acaf01c.woff",revision:"9acaf01c"},{url:"/_next/static/media/KaTeX_Size3-Regular.a144ef58.ttf",revision:"a144ef58"},{url:"/_next/static/media/KaTeX_Size3-Regular.b4230e7e.woff2",revision:"b4230e7e"},{url:"/_next/static/media/KaTeX_Size4-Regular.10d95fd3.woff2",revision:"10d95fd3"},{url:"/_next/static/media/KaTeX_Size4-Regular.7a996c9d.woff",revision:"7a996c9d"},{url:"/_next/static/media/KaTeX_Size4-Regular.fbccdabe.ttf",revision:"fbccdabe"},{url:"/_next/static/media/KaTeX_Typewriter-Regular.6258592b.woff",revision:"6258592b"},{url:"/_next/static/media/KaTeX_Typewriter-Regular.a8709e36.woff2",revision:"a8709e36"},{url:"/_next/static/media/KaTeX_Typewriter-Regular.d97aaf4a.ttf",revision:"d97aaf4a"},{url:"/_next/static/media/a5b77b63ef20339c-s.woff2",revision:"96e992d510ed36aa573ab75df8698b42"},{url:"/_next/static/media/a6d330d7873e7320-s.woff2",revision:"f7ec4e2d6c9f82076c56a871d1d23a2d"},{url:"/_next/static/media/baf12dd90520ae41-s.woff2",revision:"8096f9b1a15c26638179b6c9499ff260"},{url:"/_next/static/media/bbdb6f0234009aba-s.woff2",revision:"5756151c819325914806c6be65088b13"},{url:"/_next/static/media/bd976642b4f7fd99-s.woff2",revision:"cc0ffafe16e997fe75c32c5c6837e781"},{url:"/_next/static/media/cff529cd86cc0276-s.woff2",revision:"c2b2c28b98016afb2cb7e029c23f1f9f"},{url:"/_next/static/media/d117eea74e01de14-s.woff2",revision:"4d1e5298f2c7e19ba39a6ac8d88e91bd"},{url:"/_next/static/media/de9eb3a9f0fa9e10-s.woff2",revision:"7155c037c22abdc74e4e6be351c0593c"},{url:"/_next/static/media/dfa8b99978df7bbc-s.woff2",revision:"7a500aa24dccfcf0cc60f781072614f5"},{url:"/_next/static/media/e25729ca87cc7df9-s.woff2",revision:"9a74bbc5f0d651f8f5b6df4fb3c5c755"},{url:"/_next/static/media/eb52b768f62eeeb4-s.woff2",revision:"90687dc5a4b6b6271c9f1c1d4986ca10"},{url:"/_next/static/media/f06116e890b3dadb-s.woff2",revision:"2855f7c90916c37fe4e6bd36205a26a8"},{url:"/_next/static/media/not-found.aa00b990.svg",revision:"ee9449c9c8b33e56556df99c902e940d"},{url:"/_next/static/media/python.fb8d4db4.png",revision:"6841951dd3623f17a3f6a880c3f4f0a0"},{url:"/android-chrome-192x192.png",revision:"84c28861078626330c631ba955b97993"},{url:"/android-chrome-512x512.png",revision:"a6ca2ee37862abe8589a757cb8045e0f"},{url:"/apple-touch-icon.png",revision:"42cb8727b1f66cfd3672f326f79a376d"},{url:"/browserconfig.xml",revision:"1f9bfebf55d1737f682ecd9caffe569e"},{url:"/favicon-16x16.png",revision:"f48151236da77c692831d587f99519e7"},{url:"/favicon-32x32.png",revision:"f4aea7ba48e0d025dccdd7f2f5d8509a"},{url:"/favicon.ico",revision:"54fb3d586ff0a1cc2774782b682aebf0"},{url:"/gradients/looper-pattern.svg",revision:"d09dd7a51f00307a9ab0ece45f798d36"},{url:"/logo.svg",revision:"1b6b90d3e401f1433701507193ab988c"},{url:"/manifest.json",revision:"84653fe4bb5dbbfc4918428541283aab"},{url:"/mstile-144x144.png",revision:"790f406bd5a28023b69f413278d45731"},{url:"/mstile-150x150.png",revision:"3dff7c026a96ae7b1f162bf182e7a4a0"},{url:"/mstile-310x150.png",revision:"92a40fb99a71fc0e548381c405b8fdf7"},{url:"/mstile-310x310.png",revision:"5385bd69179e99eb9b3304fdc53c48c8"},{url:"/mstile-70x70.png",revision:"0298f1436f55f7a49179290a9aea10cd"},{url:"/robots.txt",revision:"a6850f1716e9261d1f8dca093eca3392"},{url:"/safari-pinned-tab.svg",revision:"dc12996163aa8dcfcbf450f7767a86d8"},{url:"/search-meta.json",revision:"b811bc6613a7b6a77976716336d72440"},{url:"/sitemap-0.xml",revision:"c033ad8aab8a821111ca99264bb64318"},{url:"/sitemap.xml",revision:"ebbca34a7cd10f5467aecf3331106bcf"}],{ignoreURLParametersMatching:[]}),e.cleanupOutdatedCaches(),e.registerRoute("/",new e.NetworkFirst({cacheName:"start-url",plugins:[{cacheWillUpdate:async({request:e,response:a,event:i,state:t})=>a&&"opaqueredirect"===a.type?new Response(a.body,{status:200,statusText:"OK",headers:a.headers}):a}]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,new e.CacheFirst({cacheName:"google-fonts-webfonts",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:31536e3})]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,new e.StaleWhileRevalidate({cacheName:"google-fonts-stylesheets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,new e.StaleWhileRevalidate({cacheName:"static-font-assets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,new e.StaleWhileRevalidate({cacheName:"static-image-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/image\?url=.+$/i,new e.StaleWhileRevalidate({cacheName:"next-image",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp3|wav|ogg)$/i,new e.CacheFirst({cacheName:"static-audio-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp4)$/i,new e.CacheFirst({cacheName:"static-video-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:js)$/i,new e.StaleWhileRevalidate({cacheName:"static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:css|less)$/i,new e.StaleWhileRevalidate({cacheName:"static-style-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/data\/.+\/.+\.json$/i,new e.StaleWhileRevalidate({cacheName:"next-data",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:json|xml|csv)$/i,new e.NetworkFirst({cacheName:"static-data-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>{if(!(self.origin===e.origin))return!1;const a=e.pathname;return!a.startsWith("/api/auth/")&&!!a.startsWith("/api/")}),new e.NetworkFirst({cacheName:"apis",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:16,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>{if(!(self.origin===e.origin))return!1;return!e.pathname.startsWith("/api/")}),new e.NetworkFirst({cacheName:"others",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>!(self.origin===e.origin)),new e.NetworkFirst({cacheName:"cross-origin",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:3600})]}),"GET")}));
