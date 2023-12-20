package main

func getInput01() string {
	return `%pj -> sh
%mn -> jp
&hf -> rg, vl, tq, qq, mv, zz
%xl -> hf
%sv -> mn, dl
%kk -> lh
&sj -> kz
%jj -> lq, kk
%mr -> bm, hb
%sx -> lq, fn
%fn -> zr, lq
%pf -> dl, gv
%lr -> jj, lq
%jp -> dl, pj
&hb -> sj, mr, rz, qg, pr
%vg -> zz, hf
%pr -> zq
%hn -> pf
%jg -> tj
%qg -> vk
%dv -> xl, hf
&qq -> kz
%fm -> lr
&ls -> kz
%pd -> hb, xg
%rj -> hb
%fb -> hf, tq
%zz -> np
%bm -> pd, hb
%xn -> lq, fm
%gv -> jg, dl
%dz -> sx
%zs -> dl, nh
%tj -> zs, dl
%mv -> vl
&kz -> rx
%np -> cl, hf
&bg -> kz
%vl -> vg
%xg -> rz, hb
%rz -> pr
%zq -> hb, qg
%lh -> rd
%zr -> lq
%fl -> hb, rj
%xr -> xn, lq
%rd -> dz, lq
%cl -> hf, gj
%nh -> dl
%sh -> hn, dl
%vk -> fx, hb
%gj -> hf, dv
%tq -> mv
&dl -> hn, pj, ls, mn, jg, sv
%fx -> fl, hb
&lq -> bg, kk, dz, xr, lh, fm
%rg -> hf, fb
broadcaster -> xr, mr, rg, sv`
}

func getTestInput01() string {
	return `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`
}

func getTestInput02() string {
	return `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`
}

func getInput() string {
	return getInput01()
}
