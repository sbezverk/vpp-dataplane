#!/bin/bash
VPP_COMMIT=d20bc1d3

if [ ! -d $1 ]; then
	git clone "https://gerrit.fd.io/r/vpp" $1
	cd $1
	git reset --hard ${VPP_COMMIT}
else
	cd $1
	git fetch "https://gerrit.fd.io/r/vpp" && git reset --hard ${VPP_COMMIT}
fi

git fetch "https://gerrit.fd.io/r/vpp" refs/changes/11/28711/4 && git cherry-pick FETCH_HEAD # vlib: force input node interrupts to be unique
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/87/28587/20 && git cherry-pick FETCH_HEAD # calico plugin
# Policies
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/83/28083/10 && git cherry-pick FETCH_HEAD # ACL custom policies
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/13/28513/9 && git cherry-pick FETCH_HEAD # Calico policies
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/15/29215/4 && git cherry-pick FETCH_HEAD # Calico matching
