```
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596154730771_70961
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Get:1 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:4 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Ign:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:11 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:12 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:14 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Get:16 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Fetched 7,789 kB in 20s (378 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp_7_7v1b_/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp_7_7v1b_/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmp_7_7v1b_/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Ign:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Ign:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Ign:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Err:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1
    default:   400  Bad Request [IP: 91.189.88.142 80]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Fetched 6,875 kB in 24s (283 kB/s)
    default: E
    default: : 
    default: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/p/python-cffi/python-cffi-backend_1.5.2-1ubuntu1_amd64.deb  400  Bad Request [IP: 91.189.88.142 80]
    default: E
    default: : 
    default: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 58.1 kB/6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 58.1 kB in 0s (58.4 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: [WARNING]: The value True (type bool) in a string field was converted to
    default: u'True' (type string). If this does not look like what you expect, quote the
    default: entire value to ensure it does not change.
    default: [WARNING]: Updating cache and auto-installing missing dependency: python-apt
    default: changed: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "apt cache update failed"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=5    changed=4    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596155884494_49153
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:5 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Ign:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:10 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:12 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:13 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:14 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Fetched 7,789 kB in 23s (335 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpegmetj6e/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpegmetj6e/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpegmetj6e/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Ign:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Ign:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Ign:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Err:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1
    default:   400  Bad Request [IP: 91.189.88.152 80]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Fetched 6,875 kB in 22s (305 kB/s)
    default: E
    default: : 
    default: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/p/python-cffi/python-cffi-backend_1.5.2-1ubuntu1_amd64.deb  400  Bad Request [IP: 91.189.88.152 80]
    default: E
    default: : 
    default: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 58.1 kB/6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 58.1 kB in 0s (58.7 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: [WARNING]: The value True (type bool) in a string field was converted to
    default: u'True' (type string). If this does not look like what you expect, quote the
    default: entire value to ensure it does not change.
    default: [WARNING]: Updating cache and auto-installing missing dependency: python-apt
    default: changed: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "apt cache update failed"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=6    changed=4    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo apt-get update
[sudo] nykim의 암호: 
받기:1 https://download.docker.com/linux/ubuntu bionic InRelease [64.4 kB]
기존:2 http://dl.google.com/linux/chrome/deb stable InRelease                  
기존:3 https://packages.microsoft.com/repos/vscode stable InRelease            
기존:4 https://typora.io/linux ./ InRelease                                    
기존:5 http://ppa.launchpad.net/ansible/ansible/ubuntu bionic InRelease        
기존:6 http://kr.archive.ubuntu.com/ubuntu bionic InRelease                    
받기:7 http://kr.archive.ubuntu.com/ubuntu bionic-updates InRelease [88.7 kB]  
받기:8 http://security.ubuntu.com/ubuntu bionic-security InRelease [88.7 kB]   
기존:9 http://ppa.launchpad.net/linuxgndu/sqlitebrowser/ubuntu bionic InRelease
받기:10 http://kr.archive.ubuntu.com/ubuntu bionic-backports InRelease [74.6 kB]
내려받기 316 k바이트, 소요시간 3초 (120 k바이트/초)                
패키지 목록을 읽는 중입니다... 완료
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo apt-get upgrade
패키지 목록을 읽는 중입니다... 완료
의존성 트리를 만드는 중입니다       
상태 정보를 읽는 중입니다... 완료
업그레이드를 계산하는 중입니다... 완료
0개 업그레이드, 0개 새로 설치, 0개 제거 및 0개 업그레이드 안 함.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596156963048_68728
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Get:1 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:4 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Ign:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:11 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:12 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:13 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:15 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Fetched 7,789 kB in 23s (330 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpzgl6s0ge/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpzgl6s0ge/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpzgl6s0ge/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Ign:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Ign:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Ign:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Err:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1
    default:   400  Bad Request [IP: 91.189.88.142 80]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Fetched 6,875 kB in 23s (291 kB/s)
    default: E
    default: : 
    default: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/p/python-cffi/python-cffi-backend_1.5.2-1ubuntu1_amd64.deb  400  Bad Request [IP: 91.189.88.142 80]
    default: E
    default: : 
    default: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 58.1 kB/6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 58.1 kB in 1s (54.4 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: ERROR! We were unable to read either as JSON nor YAML, these are the errors we got from each:
    default: JSON: No JSON object could be decoded
    default: 
    default: Syntax Error while loading YAML.
    default:   mapping values are not allowed in this context
    default: 
    default: The error appears to be in '/vagrant/ansible/init/tasks/main.yml': line 14, column 20, but may
    default: be elsewhere in the file depending on the exact syntax problem.
    default: 
    default: The offending line appears to be:
    default: 
    default:   apt: update_cache=yes
    default:                    ^ here
    default: 
    default: There appears to be both 'k=v' shorthand syntax and YAML in this task. Only one syntax may be used.
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596157171403_95945
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Get:1 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Get:2 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Get:8 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:9 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Ign:10 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:12 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:14 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Fetched 7,789 kB in 33s (234 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmppr6tlt27/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmppr6tlt27/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmppr6tlt27/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Ign:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Ign:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Ign:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1 [108 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 6,934 kB in 32s (213 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: ERROR! conflicting action statements: apt, stderr
    default: 
    default: The error appears to be in '/vagrant/ansible/init/tasks/main.yml': line 12, column 3, but may
    default: be elsewhere in the file depending on the exact syntax problem.
    default: 
    default: The offending line appears to be:
    default: 
    default: 
    default: - name: Update and upgrade apt packages
    default:   ^ here
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpwng_sdzb/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpwng_sdzb/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpwng_sdzb/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: [WARNING]: Updating cache and auto-installing missing dependency: python-apt
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "'/usr/bin/apt-get upgrade --with-new-pkgs ' failed: E: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/g/grub2/grub-pc_2.02~beta2-36ubuntu3.26_amd64.deb  400  Bad Request [IP: 91.189.88.142 80]\n\nE: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/g/grub2/grub-common_2.02~beta2-36ubuntu3.26_amd64.deb  400  Bad Request [IP: 91.189.88.142 80]\n\nE: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?\n", "rc": 100, "stdout": "Reading package lists...\nBuilding dependency tree...\nReading state information...\nCalculating upgrade...\nThe following NEW packages will be installed:\n  linux-image-4.4.0-186-generic linux-modules-4.4.0-186-generic\n  linux-modules-extra-4.4.0-186-generic\nThe following packages will be upgraded:\n  grub-common grub-pc grub-pc-bin grub2-common libpython2.7\n  libpython2.7-minimal libpython2.7-stdlib libpython3.5 libpython3.5-minimal\n  libpython3.5-stdlib libseccomp2 linux-image-generic linux-libc-dev python2.7\n  python2.7-minimal python3.5 python3.5-minimal sosreport\n18 upgraded, 3 newly installed, 0 to remove and 0 not upgraded.\nNeed to get 70.4 MB of archives.\nAfter this operation, 224 MB of additional disk space will be used.\nGet:1 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5 amd64 3.5.2-2ubuntu0~16.04.11 [1360 kB]\nGet:2 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5 amd64 3.5.2-2ubuntu0~16.04.11 [165 kB]\nErr:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11\n  400  Bad Request [IP: 91.189.88.142 80]\nGet:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11 [2132 kB]\nGet:4 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [1597 kB]\nGet:5 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [525 kB]\nErr:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc amd64 2.02~beta2-36ubuntu3.26\n  400  Bad Request [IP: 91.189.88.142 80]\nGet:7 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub2-common amd64 2.02~beta2-36ubuntu3.26 [510 kB]\nGet:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc-bin amd64 2.02~beta2-36ubuntu3.26 [891 kB]\nErr:9 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-common amd64 2.02~beta2-36ubuntu3.26\n  400  Bad Request [IP: 91.189.88.142 80]\nGet:10 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7 amd64 2.7.12-1ubuntu0~16.04.12 [224 kB]\nGet:11 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7 amd64 2.7.12-1ubuntu0~16.04.12 [1070 kB]\nErr:12 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython2.7-stdlib amd64 2.7.12-1ubuntu0~16.04.12\n  400  Bad Request [IP: 91.189.88.142 80]\nGet:12 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython2.7-stdlib amd64 2.7.12-1ubuntu0~16.04.12 [1883 kB]\nGet:13 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [1258 kB]\nGet:14 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [338 kB]\nGet:15 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libseccomp2 amd64 2.4.3-1ubuntu3.16.04.3 [41.0 kB]\nGet:16 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-4.4.0-186-generic amd64 4.4.0-186.216 [12.0 MB]\nGet:17 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-4.4.0-186-generic amd64 4.4.0-186.216 [6945 kB]\nGet:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-extra-4.4.0-186-generic amd64 4.4.0-186.216 [36.6 MB]\nGet:19 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-generic amd64 4.4.0.186.192 [2514 B]\nGet:20 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-libc-dev amd64 4.4.0
    default: -186.216 [853 kB]\nGet:21 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 sosreport amd64 3.9.1-1ubuntu0.16.04.1 [170 kB]\nFetched 68.5 MB in 4min 20s (263 kB/s)\n", "stdout_lines": ["Reading package lists...", "Building dependency tree...", "Reading state information...", "Calculating upgrade...", "The following NEW packages will be installed:", "  linux-image-4.4.0-186-generic linux-modules-4.4.0-186-generic", "  linux-modules-extra-4.4.0-186-generic", "The following packages will be upgraded:", "  grub-common grub-pc grub-pc-bin grub2-common libpython2.7", "  libpython2.7-minimal libpython2.7-stdlib libpython3.5 libpython3.5-minimal", "  libpython3.5-stdlib libseccomp2 linux-image-generic linux-libc-dev python2.7", "  python2.7-minimal python3.5 python3.5-minimal sosreport", "18 upgraded, 3 newly installed, 0 to remove and 0 not upgraded.", "Need to get 70.4 MB of archives.", "After this operation, 224 MB of additional disk space will be used.", "Get:1 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5 amd64 3.5.2-2ubuntu0~16.04.11 [1360 kB]", "Get:2 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5 amd64 3.5.2-2ubuntu0~16.04.11 [165 kB]", "Err:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11", "  400  Bad Request [IP: 91.189.88.142 80]", "Get:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11 [2132 kB]", "Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [1597 kB]", "Get:5 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [525 kB]", "Err:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc amd64 2.02~beta2-36ubuntu3.26", "  400  Bad Request [IP: 91.189.88.142 80]", "Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub2-common amd64 2.02~beta2-36ubuntu3.26 [510 kB]", "Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc-bin amd64 2.02~beta2-36ubuntu3.26 [891 kB]", "Err:9 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-common amd64 2.02~beta2-36ubuntu3.26", "  400  Bad Request [IP: 91.189.88.142 80]", "Get:10 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7 amd64 2.7.12-1ubuntu0~16.04.12 [224 kB]", "Get:11 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7 amd64 2.7.12-1ubuntu0~16.04.12 [1070 kB]", "Err:12 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython2.7-stdlib amd64 2.7.12-1ubuntu0~16.04.12", "  400  Bad Request [IP: 91.189.88.142 80]", "Get:12 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython2.7-stdlib amd64 2.7.12-1ubuntu0~16.04.12 [1883 kB]", "Get:13 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [1258 kB]", "Get:14 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [338 kB]", "Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libseccomp2 amd64 2.4.3-1ubuntu3.16.04.3 [41.0 kB]", "Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-4.4.0-186-generic amd64 4.4.0-186.216 [12.0 MB]", "Get:17 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-4.4.0-186-generic amd64 4.4.0-186.216 [6945 kB]", "Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-extra-4.4.0-186-generic amd64 4.4.0-186.216 [36.6 MB]", "Get:19 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-generic amd64 4.4.0.186.192 [2514 B]", "Get:20 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-libc-dev amd64 4.4.0-186.216 [853 kB]", "Get:21 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 sosreport amd64 3.9.1-1ubuntu0.16.04.1 [170 kB]", "Fetched 68.5 MB in 4min 20s (263 kB/s)"]}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=4    changed=3    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:3 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpqg7yv7jt/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpqg7yv7jt/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpqg7yv7jt/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Unsupported parameters for (apt) module: stderr, stdout Supported parameters include: allow_unauthenticated, autoclean, autoremove, cache_valid_time, deb, default_release, dpkg_options, force, force_apt_get, install_recommends, only_upgrade, package, policy_rc_d, purge, state, update_cache, upgrade"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=5    changed=1    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:2 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpdmrgy53p/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpdmrgy53p/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpdmrgy53p/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "apt cache update failed"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=7    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision

==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:3 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpp4kxqerb/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpp4kxqerb/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpp4kxqerb/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: ERROR! conflicting action statements: apt_repository, apt
    default: 
    default: The error appears to be in '/vagrant/ansible/init/tasks/main.yml': line 34, column 3, but may
    default: be elsewhere in the file depending on the exact syntax problem.
    default: 
    default: The offending line appears to be:
    default: 
    default: 
    default: - name: Add python3.6 repo
    default:   ^ here
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpzwdp3_6k/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpzwdp3_6k/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpzwdp3_6k/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Unsupported parameters for (apt_repository) module: upgrade Supported parameters include: codename, filename, install_python_apt, mode, repo, state, update_cache, validate_certs"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=7    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:2 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:3 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp707eb5yt/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp707eb5yt/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmp707eb5yt/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "apt cache update failed"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=7    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision

==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Fetched 3,966 kB in 19s (202 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp977iejxl/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp977iejxl/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: /tmp/tmp977iejxl/trustdb.gpg: trustdb created
    default: gpg: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: Total number processed: 1
    default: gpg:               imported: 1  (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : update again] *****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "apt cache update failed"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=8    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Get:1 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:2 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Fetched 216 kB in 1s (133 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmptfk82wk2/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmptfk82wk2/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmptfk82wk2/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : update again] *****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "apt cache update failed"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=8    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Fetched 216 kB in 1s (132 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpwkd6hn3y/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpwkd6hn3y/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpwkd6hn3y/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add python3.6 repo] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : update again] *****************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Failed to update apt cache: W:The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file., W:Data from such a repository can't be authenticated and is therefore potentially dangerous to use., W:See apt-secure(8) manpage for repository creation and user configuration details., E:Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden, E:Some index files failed to download. They have been ignored, or old ones used instead."}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=8    changed=1    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596159788820_62580
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:2 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:4 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Ign:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:11 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:13 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:14 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Get:16 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Fetched 7,790 kB in 21s (365 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpeg5aqxqk/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpeg5aqxqk/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpeg5aqxqk/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Ign:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Ign:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Ign:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Err:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1
    default:   400  Bad Request [IP: 91.189.88.152 80]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Fetched 6,875 kB in 26s (264 kB/s)
    default: E
    default: : 
    default: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/p/python-cffi/python-cffi-backend_1.5.2-1ubuntu1_amd64.deb  400  Bad Request [IP: 91.189.88.152 80]
    default: E
    default: : 
    default: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 58.1 kB/6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 58.1 kB in 1s (58.1 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: [WARNING]: Updating cache and auto-installing missing dependency: python-apt
    default: changed: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : update again] *****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Install basic packages] *******************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "No package matching 'python3.6' is available"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=8    changed=4    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:2 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpltmyz5mo/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpltmyz5mo/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpltmyz5mo/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : add python-apt] ***************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "missing required arguments: repo"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=7    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:2 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:3 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpekpk1w2x/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpekpk1w2x/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpekpk1w2x/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : add recommend] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : update again] *****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Install basic packages] *******************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "No package matching 'python3.6' is available"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=9    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:3 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpomzlqxsr/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpomzlqxsr/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpomzlqxsr/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : add recommend] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add key for jonathonf PPA] ****************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Add jonathonf PPA] ************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "apt cache update failed"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=9    changed=1    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmprkei4dea/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmprkei4dea/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmprkei4dea/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : add recommend] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add key for jonathonf PPA] ****************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add jonathonf PPA] ************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : update again] *****************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Failed to update apt cache: W:The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file., W:Data from such a repository can't be authenticated and is therefore potentially dangerous to use., W:See apt-secure(8) manpage for repository creation and user configuration details., E:Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden, E:Some index files failed to download. They have been ignored, or old ones used instead."}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=10   changed=1    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:2 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Ign:5 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial InRelease
    default: Hit:6 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Err:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default:   403  Forbidden
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Reading package lists...
    default: W
    default: : 
    default: The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file.
    default: E
    default: : 
    default: Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden
    default: E
    default: : 
    default: Some index files failed to download. They have been ignored, or old ones used instead.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpjouyf4gk/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpjouyf4gk/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpjouyf4gk/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Failed to update apt cache: W:The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file., W:Data from such a repository can't be authenticated and is therefore potentially dangerous to use., W:See apt-secure(8) manpage for repository creation and user configuration details., E:Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden, E:Some index files failed to download. They have been ignored, or old ones used instead."}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=4    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi V
[sudo] nykim의 암호: 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi Vagrantfile 
[sudo] nykim의 암호: 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi Vagrantfile 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:2 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:3 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Ign:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial InRelease
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Err:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default:   403  Forbidden
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Reading package lists...
    default: W
    default: : 
    default: The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file.
    default: E
    default: : 
    default: Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden
    default: E
    default: : 
    default: Some index files failed to download. They have been ignored, or old ones used instead.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp7mojty06/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp7mojty06/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmp7mojty06/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Failed to update apt cache: W:The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file., W:Data from such a repository can't be authenticated and is therefore potentially dangerous to use., W:See apt-secure(8) manpage for repository creation and user configuration details., E:Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden, E:Some index files failed to download. They have been ignored, or old ones used instead."}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=4    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Ign:2 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Ign:4 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Ign:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Get:8 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Err:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default:   403  Forbidden
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Fetched 325 kB in 25s (12.9 kB/s)
    default: Reading package lists...
    default: W
    default: : 
    default: The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file.
    default: E
    default: : 
    default: Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden
    default: E
    default: : 
    default: Some index files failed to download. They have been ignored, or old ones used instead.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpwx_xhpkn/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpwx_xhpkn/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpwx_xhpkn/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Failed to update apt cache: W:The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file., W:Data from such a repository can't be authenticated and is therefore potentially dangerous to use., W:See apt-secure(8) manpage for repository creation and user configuration details., E:Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden, E:Some index files failed to download. They have been ignored, or old ones used instead."}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=4    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:4 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Ign:6 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial InRelease
    default: Ign:7 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Ign:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Err:8 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main amd64 Packages
    default:   403  Forbidden
    default: Ign:9 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main i386 Packages
    default: Ign:10 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main all Packages
    default: Ign:11 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en_US
    default: Ign:12 http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial/main Translation-en
    default: Reading package lists...
    default: W
    default: : 
    default: The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file.
    default: E
    default: : 
    default: Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden
    default: E
    default: : 
    default: Some index files failed to download. They have been ignored, or old ones used instead.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmptagciwbg/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmptagciwbg/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmptagciwbg/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Failed to update apt cache: W:The repository 'http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu xenial Release' does not have a Release file., W:Data from such a repository can't be authenticated and is therefore potentially dangerous to use., W:See apt-secure(8) manpage for repository creation and user configuration details., E:Failed to fetch http://ppa.launchpad.net/jonathonf/python-3.6/ubuntu/dists/xenial/main/binary-amd64/Packages  403  Forbidden, E:Some index files failed to download. They have been ignored, or old ones used instead."}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=4    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596168502309_17841
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Get:1 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:4 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Get:8 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:9 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:10 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Get:12 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Fetched 7,790 kB in 2min 27s (52.7 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp09dw7fe4/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp09dw7fe4/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: /tmp/tmp09dw7fe4/trustdb.gpg: trustdb created
    default: gpg: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: Total number processed: 1
    default: gpg:               imported: 1  (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 6,934 kB in 1min 16s (91.1 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
^C==> default: Waiting for cleanup before exiting...
^C==> default: Exiting immediately, without cleanup!
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596169126511_39507
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Connection reset. Retrying...
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Get:1 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:4 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Ign:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:11 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:12 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:13 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:15 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Fetched 7,790 kB in 23s (333 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp0hkrbwoo/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp0hkrbwoo/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmp0hkrbwoo/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Ign:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Ign:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Ign:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Err:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1
    default:   400  Bad Request [IP: 91.189.88.142 80]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main i386 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Fetched 6,875 kB in 22s (303 kB/s)
    default: E
    default: : 
    default: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/p/python-cffi/python-cffi-backend_1.5.2-1ubuntu1_amd64.deb  400  Bad Request [IP: 91.189.88.142 80]
    default: E
    default: : 
    default: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 58.1 kB/6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 58.1 kB in 1s (44.9 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Install requires package for apt module] **************************
    default: ok: [localhost]
    default: [WARNING]: Updating cache and auto-installing missing dependency: python-apt
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "'/usr/bin/apt-get upgrade --with-new-pkgs ' failed: E: Failed to fetch http://archive.ubuntu.com/ubuntu/pool/main/g/grub2/grub-pc_2.02~beta2-36ubuntu3.26_amd64.deb  400  Bad Request [IP: 91.189.88.152 80]\n\nE: Unable to fetch some archives, maybe run apt-get update or try with --fix-missing?\n", "rc": 100, "stdout": "Reading package lists...\nBuilding dependency tree...\nReading state information...\nCalculating upgrade...\nThe following NEW packages will be installed:\n  linux-image-4.4.0-186-generic linux-modules-4.4.0-186-generic\n  linux-modules-extra-4.4.0-186-generic\nThe following packages will be upgraded:\n  grub-common grub-pc grub-pc-bin grub2-common libpython2.7\n  libpython2.7-minimal libpython2.7-stdlib libpython3.5 libpython3.5-minimal\n  libpython3.5-stdlib libseccomp2 linux-image-generic linux-libc-dev python2.7\n  python2.7-minimal python3.5 python3.5-minimal sosreport\n18 upgraded, 3 newly installed, 0 to remove and 0 not upgraded.\nNeed to get 70.4 MB of archives.\nAfter this operation, 224 MB of additional disk space will be used.\nGet:1 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5 amd64 3.5.2-2ubuntu0~16.04.11 [1360 kB]\nGet:2 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5 amd64 3.5.2-2ubuntu0~16.04.11 [165 kB]\nErr:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11\n  400  Bad Request [IP: 91.189.88.152 80]\nGet:4 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [1597 kB]\nGet:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11 [2132 kB]\nGet:5 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [525 kB]\nErr:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc amd64 2.02~beta2-36ubuntu3.26\n  400  Bad Request [IP: 91.189.88.152 80]\nGet:7 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc-bin amd64 2.02~beta2-36ubuntu3.26 [891 kB]\nGet:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub2-common amd64 2.02~beta2-36ubuntu3.26 [510 kB]\nGet:9 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-common amd64 2.02~beta2-36ubuntu3.26 [1704 kB]\nGet:10 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7 amd64 2.7.12-1ubuntu0~16.04.12 [1070 kB]\nGet:11 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7 amd64 2.7.12-1ubuntu0~16.04.12 [224 kB]\nGet:12 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7-stdlib amd64 2.7.12-1ubuntu0~16.04.12 [1883 kB]\nGet:13 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [1258 kB]\nGet:14 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [338 kB]\nGet:15 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libseccomp2 amd64 2.4.3-1ubuntu3.16.04.3 [41.0 kB]\nGet:16 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-4.4.0-186-generic amd64 4.4.0-186.216 [12.0 MB]\nGet:17 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-4.4.0-186-generic amd64 4.4.0-186.216 [6945 kB]\nGet:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-extra-4.4.0-186-generic amd64 4.4.0-186.216 [36.6 MB]\nGet:19 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-generic amd64 4.4.0.186.192 [2514 B]\nGet:20 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-libc-dev amd64 4.4.0-186.216 [853 kB]\nGet:21 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 sosreport amd64 3.9.1-1ubuntu0.16.04.1 [170 kB]\nFetched 70.2 MB in 4min 44s (247 kB/s)\n", "stdout_lines": ["Reading package lists...", "Building dependency tree...", "Reading state information...", "Calculating upgrade...", "The following NEW packages will be installed:", "  linux-image-4.4.0-186-generic linux-modules-4.4.0-186-generic", "  linux-modules-extra-4.4.0-186-generic", "The following packages will be upgraded:", "  grub-common grub-pc grub-pc-bin grub2-common libpython2.7", "  libpython2.7-minimal libpython2.7-stdlib libpython3.5 libpython3.5-minimal", "  libpython3.5-stdlib libseccomp2 linux-image-generic linux-libc-dev python2.7", "  python2.7-minimal python3.5 python3.5-minimal sosreport", "18 upgraded, 3 newly installed, 0 to remove and 0 not upgraded.", "Need to get 70.4 MB of archives.", "After this operation, 224 MB of additional disk space will be used.", "Get:1 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5 amd64 3.5.2-2ubuntu0~16.04.11 [1360 kB]", "Get:2 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5 amd64 3.5.2-2ubuntu0~16.04.11 [165 kB]", "Err:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11", "  400  Bad Request [IP: 91.189.88.152 80]", "Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [1597 kB]", "Get:3 http://security.ubuntu.com/ubuntu xenial-security/main amd64 libpython3.5-stdlib amd64 3.5.2-2ubuntu0~16.04.11 [2132 kB]", "Get:5 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython3.5-minimal amd64 3.5.2-2ubuntu0~16.04.11 [525 kB]", "Err:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc amd64 2.02~beta2-36ubuntu3.26", "  400  Bad Request [IP: 91.189.88.152 80]", "Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-pc-bin amd64 2.02~beta2-36ubuntu3.26 [891 kB]", "Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub2-common amd64 2.02~beta2-36ubuntu3.26 [510 kB]", "Get:9 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 grub-common amd64 2.02~beta2-36ubuntu3.26 [1704 kB]", "Get:10 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7 amd64 2.7.12-1ubuntu0~16.04.12 [1070 kB]", "Get:11 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7 amd64 2.7.12-1ubuntu0~16.04.12 [224 kB]", "Get:12 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7-stdlib amd64 2.7.12-1ubuntu0~16.04.12 [1883 kB]", "Get:13 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [1258 kB]", "Get:14 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libpython2.7-minimal amd64 2.7.12-1ubuntu0~16.04.12 [338 kB]", "Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 libseccomp2 amd64 2.4.3-1ubuntu3.16.04.3 [41.0 kB]", "Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-4.4.0-186-generic amd64 4.4.0-186.216 [12.0 MB]", "Get:17 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-4.4.0-186-generic amd64 4.4.0-186.216 [6945 kB]", "Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-modules-extra-4.4.0-186-generic amd64 4.4.0-186.216 [36.6 MB]", "Get:19 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-image-generic amd64 4.4.0.186.192 [2514 B]", "Get:20 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 linux-libc-dev amd64 4.4.0-186.216 [853 kB]", "Get:21 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 sosreport amd64 3.9.1-1ubuntu0.16.04.1 [170 kB]", "Fetched 70.2 MB in 4min 44s (247 kB/s)"]}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=5    changed=3    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi Vagrantfile 
[sudo] nykim의 암호: 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi Vagrantfile 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596169764345_99278
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Get:4 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Get:8 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:9 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:10 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:11 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Fetched 7,790 kB in 2min 4s (62.6 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpvw5bxyfm/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpvw5bxyfm/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpvw5bxyfm/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Ign:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 6,934 kB in 1min 12s (95.4 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default:  This PPA contains more recent Python versions packaged for Ubuntu.
    default: 
    default: Disclaimer: there's no guarantee of timely updates in case of security problems or other issues. If you want to use them in a security-or-otherwise-critical environment (say, on a production server), you do so at your own risk.
    default: 
    default: Update Note
    default: ===========
    default: Please use this repository instead of ppa:fkrull/deadsnakes.
    default: 
    default: Reporting Issues
    default: ================
    default: 
    default: Issues can be reported in the master issue tracker at:
    default: https://github.com/deadsnakes/issues/issues
    default: 
    default: Supported Ubuntu and Python Versions
    default: ====================================
    default: 
    default: - Ubuntu 16.04 (xenial) Python 2.3 - Python 2.6, Python 3.1 - Python3.4, Python 3.6 - Python3.9
    default: - Ubuntu 18.04 (bionic) Python2.3 - Python 2.6, Python 3.1 - Python 3.5, Python3.7 - Python3.9
    default: - Ubuntu 20.04 (focal) Python3.5 - Python3.7, Python3.9
    default: - Note: Python2.7 (all), Python 3.5 (xenial), Python 3.6 (bionic), Python 3.8 (focal) are not provided by deadsnakes as upstream ubuntu provides those packages.
    default: - Note: for focal, older python versions require libssl1.0.x so they are not currently built
    default: 
    default: The packages may also work on other versions of Ubuntu or Debian, but that is not tested or supported.
    default: 
    default: Packages
    default: ========
    default: 
    default: The packages provided here are loosely based on the debian upstream packages with some modifications to make them more usable as non-default pythons and on ubuntu.  As such, the packages follow debian's patterns and often do not include a full python distribution with just `apt install python#.#`.  Here is a list of packages that may be useful along with the default install:
    default: 
    default: - `python#.#-dev`: includes development headers for building C extensions
    default: - `python#.#-venv`: provides the standard library `venv` module
    default: - `python#.#-distutils`: provides the standard library `distutils` module
    default: - `python#.#-lib2to3`: provides the `2to3-#.#` utility as well as the standard library `lib2to3` module
    default: - `python#.#-gdbm`: provides the standard library `dbm.gnu` module
    default: - `python#.#-tk`: provides the standard library `tkinter` module
    default: 
    default: Third-Party Python Modules
    default: ==========================
    default: 
    default: Python modules in the official Ubuntu repositories are packaged to work with the Python interpreters from the official repositories. Accordingly, they generally won't work with the Python interpreters from this PPA. As an exception, pure-Python modules for Python 3 will work, but any compiled extension modules won't.
    default: 
    default: To install 3rd-party Python modules, you should use the common Python packaging tools.  For an introduction into the Python packaging ecosystem and its tools, refer to the Python Packaging User Guide:
    default: https://packaging.python.org/installing/
    default: 
    default: Sources
    default: =======
    default: The package sources are available at:
    default: https://github.com/deadsnakes/
    default: 
    default: Nightly Builds
    default: ==============
    default: 
    default: For nightly builds, see ppa:deadsnakes/nightly https://launchpad.net/~deadsnakes/+archive/ubuntu/nightly
    default:  More info: https://launchpad.net/~deadsnakes/+archive/ubuntu/ppa
    default: gpg: 
    default: keyring `/tmp/tmps6qu5ads/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmps6qu5ads/pubring.gpg' created
    default: gpg: 
    default: requesting key 6A755776 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmps6qu5ads/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 6A755776: public key "Launchpad PPA for deadsnakes" imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: WARNING: 
    default: apt
    default:  
    default: does not have a stable CLI interface. 
    default: Use with caution in scripts.
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Get:2 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Get:5 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial InRelease [18.0 kB]
    default: Hit:6 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Get:7 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial/main amd64 Packages [31.3 kB]
    default: Get:8 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial/main i386 Packages [31.3 kB]
    default: Get:9 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial/main Translation-en [7,088 B]
    default: Fetched 197 kB in 5s (36.4 kB/s)
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: 18 packages can be upgraded. Run 'apt list --upgradable' to see them.
    default: WARNING: 
    default: apt
    default:  
    default: does not have a stable CLI interface. 
    default: Use with caution in scripts.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libpython3.6-minimal libpython3.6-stdlib python3.6-minimal
    default: Suggested packages:
    default:   python3.6-venv python3.6-doc binfmt-support
    default: The following NEW packages will be installed:
    default:   libpython3.6-minimal libpython3.6-stdlib python3.6 python3.6-minimal
    default: 0 upgraded, 4 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 4,759 kB of archives.
    default: After this operation, 24.3 MB of additional disk space will be used.
    default: Do you want to continue?
    default:  [Y/n] 
    default: Abort.
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
^C==> default: Waiting for cleanup before exiting...
^C==> default: Exiting immediately, without cleanup!
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi Vagrantfile 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant up

Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-16.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-16.04' version '202007.17.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1596170370921_9881
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Connection reset. Retrying...
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /home/nykim/바탕화면/vagrant
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Get:4 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [904 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1,180 kB]
    default: Get:7 http://security.ubuntu.com/ubuntu xenial-security/main i386 Packages [681 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main i386 Packages [937 kB]
    default: Get:9 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [336 kB]
    default: Get:10 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [497 kB]
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [442 kB]
    default: Get:12 http://security.ubuntu.com/ubuntu xenial-security/universe i386 Packages [425 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [801 kB]
    default: Get:14 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [204 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial-updates/universe i386 Packages [723 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [335 kB]
    default: Fetched 7,790 kB in 31s (244 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp51hpy6bh/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp51hpy6bh/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmp51hpy6bh/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: Suggested packages:
    default:   python-cryptography-doc python-cryptography-vectors python-enum34-doc
    default:   python-jinja2-doc doc-base python-setuptools-doc
    default: The following NEW packages will be installed:
    default:   ansible libyaml-0-2 python-cffi-backend python-cryptography python-ecdsa
    default:   python-enum34 python-httplib2 python-idna python-ipaddress python-jinja2
    default:   python-markupsafe python-paramiko python-pkg-resources python-pyasn1
    default:   python-setuptools python-six python-yaml sshpass
    default: 0 upgraded, 18 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 6,934 kB of archives.
    default: After this operation, 64.2 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial/main amd64 ansible all 2.9.11-1ppa~xenial [5,790 kB]
    default: Get:2 http://archive.ubuntu.com/ubuntu xenial/main amd64 libyaml-0-2 amd64 0.1.6-3 [47.6 kB]
    default: Get:3 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-markupsafe amd64 0.23-2build2 [15.5 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-jinja2 all 2.8-1ubuntu0.1 [106 kB]
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-yaml amd64 3.11-3build1 [105 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-six all 1.10.0-3 [10.9 kB]
    default: Get:7 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-ecdsa all 0.13-2ubuntu0.16.04.1 [36.2 kB]
    default: Get:8 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-paramiko all 1.16.0-1ubuntu0.2 [110 kB]
    default: Get:9 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-httplib2 all 0.9.1+dfsg-1 [34.2 kB]
    default: Get:10 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pkg-resources all 20.7.0-1 [108 kB]
    default: Get:11 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-setuptools all 20.7.0-1 [169 kB]
    default: Get:12 http://archive.ubuntu.com/ubuntu xenial/universe amd64 sshpass amd64 1.05-1 [10.5 kB]
    default: Get:13 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-cffi-backend amd64 1.5.2-1ubuntu1 [58.1 kB]
    default: Get:14 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-enum34 all 1.1.2-1 [35.8 kB]
    default: Get:15 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-idna all 2.0-3 [35.1 kB]
    default: Get:16 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-ipaddress all 1.0.16-1 [18.0 kB]
    default: Get:17 http://archive.ubuntu.com/ubuntu xenial/main amd64 python-pyasn1 all 0.1.9-1 [45.1 kB]
    default: Get:18 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 python-cryptography amd64 1.2.3-1ubuntu0.2 [199 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 6,934 kB in 47s (146 kB/s)
    default: Selecting previously unselected package libyaml-0-2:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 38580 files and directories currently installed.)
    default: Preparing to unpack .../libyaml-0-2_0.1.6-3_amd64.deb ...
    default: Unpacking libyaml-0-2:amd64 (0.1.6-3) ...
    default: Selecting previously unselected package python-markupsafe.
    default: Preparing to unpack .../python-markupsafe_0.23-2build2_amd64.deb ...
    default: Unpacking python-markupsafe (0.23-2build2) ...
    default: Selecting previously unselected package python-jinja2.
    default: Preparing to unpack .../python-jinja2_2.8-1ubuntu0.1_all.deb ...
    default: Unpacking python-jinja2 (2.8-1ubuntu0.1) ...
    default: Selecting previously unselected package python-yaml.
    default: Preparing to unpack .../python-yaml_3.11-3build1_amd64.deb ...
    default: Unpacking python-yaml (3.11-3build1) ...
    default: Selecting previously unselected package python-six.
    default: Preparing to unpack .../python-six_1.10.0-3_all.deb ...
    default: Unpacking python-six (1.10.0-3) ...
    default: Selecting previously unselected package python-ecdsa.
    default: Preparing to unpack .../python-ecdsa_0.13-2ubuntu0.16.04.1_all.deb ...
    default: Unpacking python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Selecting previously unselected package python-paramiko.
    default: Preparing to unpack .../python-paramiko_1.16.0-1ubuntu0.2_all.deb ...
    default: Unpacking python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Selecting previously unselected package python-httplib2.
    default: Preparing to unpack .../python-httplib2_0.9.1+dfsg-1_all.deb ...
    default: Unpacking python-httplib2 (0.9.1+dfsg-1) ...
    default: Selecting previously unselected package python-pkg-resources.
    default: Preparing to unpack .../python-pkg-resources_20.7.0-1_all.deb ...
    default: Unpacking python-pkg-resources (20.7.0-1) ...
    default: Selecting previously unselected package python-setuptools.
    default: Preparing to unpack .../python-setuptools_20.7.0-1_all.deb ...
    default: Unpacking python-setuptools (20.7.0-1) ...
    default: Selecting previously unselected package sshpass.
    default: Preparing to unpack .../sshpass_1.05-1_amd64.deb ...
    default: Unpacking sshpass (1.05-1) ...
    default: Selecting previously unselected package python-cffi-backend.
    default: Preparing to unpack .../python-cffi-backend_1.5.2-1ubuntu1_amd64.deb ...
    default: Unpacking python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Selecting previously unselected package python-enum34.
    default: Preparing to unpack .../python-enum34_1.1.2-1_all.deb ...
    default: Unpacking python-enum34 (1.1.2-1) ...
    default: Selecting previously unselected package python-idna.
    default: Preparing to unpack .../python-idna_2.0-3_all.deb ...
    default: Unpacking python-idna (2.0-3) ...
    default: Selecting previously unselected package python-ipaddress.
    default: Preparing to unpack .../python-ipaddress_1.0.16-1_all.deb ...
    default: Unpacking python-ipaddress (1.0.16-1) ...
    default: Selecting previously unselected package python-pyasn1.
    default: Preparing to unpack .../python-pyasn1_0.1.9-1_all.deb ...
    default: Unpacking python-pyasn1 (0.1.9-1) ...
    default: Selecting previously unselected package python-cryptography.
    default: Preparing to unpack .../python-cryptography_1.2.3-1ubuntu0.2_amd64.deb ...
    default: Unpacking python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Selecting previously unselected package ansible.
    default: Preparing to unpack .../ansible_2.9.11-1ppa~xenial_all.deb ...
    default: Unpacking ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Setting up libyaml-0-2:amd64 (0.1.6-3) ...
    default: Setting up python-markupsafe (0.23-2build2) ...
    default: Setting up python-jinja2 (2.8-1ubuntu0.1) ...
    default: Setting up python-yaml (3.11-3build1) ...
    default: Setting up python-six (1.10.0-3) ...
    default: Setting up python-ecdsa (0.13-2ubuntu0.16.04.1) ...
    default: Setting up python-paramiko (1.16.0-1ubuntu0.2) ...
    default: Setting up python-httplib2 (0.9.1+dfsg-1) ...
    default: Setting up python-pkg-resources (20.7.0-1) ...
    default: Setting up python-setuptools (20.7.0-1) ...
    default: Setting up sshpass (1.05-1) ...
    default: Setting up python-cffi-backend (1.5.2-1ubuntu1) ...
    default: Setting up python-enum34 (1.1.2-1) ...
    default: Setting up python-idna (2.0-3) ...
    default: Setting up python-ipaddress (1.0.16-1) ...
    default: Setting up python-pyasn1 (0.1.9-1) ...
    default: Setting up python-cryptography (1.2.3-1ubuntu0.2) ...
    default: Setting up ansible (2.9.11-1ppa~xenial) ...
    default: Processing triggers for libc-bin (2.23-0ubuntu11.2) ...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 18 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp_lwtc0no/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp_lwtc0no/pubring.gpg' created
    default: gpg: 
    default: requesting key 6A755776 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmp_lwtc0no/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 6A755776: public key "Launchpad PPA for deadsnakes" imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: WARNING: 
    default: apt
    default:  
    default: does not have a stable CLI interface. 
    default: Use with caution in scripts.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: The following additional packages will be installed:
    default:   libpython3.6-minimal libpython3.6-stdlib python3.6-minimal
    default: Suggested packages:
    default:   python3.6-venv python3.6-doc binfmt-support
    default: The following NEW packages will be installed:
    default:   libpython3.6-minimal libpython3.6-stdlib python3.6 python3.6-minimal
    default: 0 upgraded, 4 newly installed, 0 to remove and 18 not upgraded.
    default: Need to get 4,759 kB of archives.
    default: After this operation, 24.3 MB of additional disk space will be used.
    default: Get:1 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial/main amd64 libpython3.6-minimal amd64 3.6.11-1+xenial1 [578 kB]
    default: Get:2 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial/main amd64 python3.6-minimal amd64 3.6.11-1+xenial1 [1,714 kB]
    default: Get:3 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial/main amd64 libpython3.6-stdlib amd64 3.6.11-1+xenial1 [2,220 kB]
    default: Get:4 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial/main amd64 python3.6 amd64 3.6.11-1+xenial1 [247 kB]
    default: dpkg-preconfigure: unable to re-open stdin: No such file or directory
    default: Fetched 4,759 kB in 43s (109 kB/s)
    default: Selecting previously unselected package libpython3.6-minimal:amd64.
    default: (Reading database ... 
    default: (Reading database ... 5%
    default: (Reading database ... 10%
    default: (Reading database ... 15%
    default: (Reading database ... 20%
    default: (Reading database ... 25%
    default: (Reading database ... 30%
    default: (Reading database ... 35%
    default: (Reading database ... 40%
    default: (Reading database ... 45%
    default: (Reading database ... 50%
    default: (Reading database ... 55%
    default: (Reading database ... 60%
    default: (Reading database ... 65%
    default: (Reading database ... 70%
    default: (Reading database ... 75%
    default: (Reading database ... 80%
    default: (Reading database ... 85%
    default: (Reading database ... 90%
    default: (Reading database ... 95%
    default: (Reading database ... 100%
    default: (Reading database ... 
    default: 45167 files and directories currently installed.)
    default: Preparing to unpack .../libpython3.6-minimal_3.6.11-1+xenial1_amd64.deb ...
    default: Unpacking libpython3.6-minimal:amd64 (3.6.11-1+xenial1) ...
    default: Selecting previously unselected package python3.6-minimal.
    default: Preparing to unpack .../python3.6-minimal_3.6.11-1+xenial1_amd64.deb ...
    default: Unpacking python3.6-minimal (3.6.11-1+xenial1) ...
    default: Selecting previously unselected package libpython3.6-stdlib:amd64.
    default: Preparing to unpack .../libpython3.6-stdlib_3.6.11-1+xenial1_amd64.deb ...
    default: Unpacking libpython3.6-stdlib:amd64 (3.6.11-1+xenial1) ...
    default: Selecting previously unselected package python3.6.
    default: Preparing to unpack .../python3.6_3.6.11-1+xenial1_amd64.deb ...
    default: Unpacking python3.6 (3.6.11-1+xenial1) ...
    default: Processing triggers for man-db (2.7.5-1) ...
    default: Processing triggers for mime-support (3.59ubuntu1) ...
    default: Setting up libpython3.6-minimal:amd64 (3.6.11-1+xenial1) ...
    default: Setting up python3.6-minimal (3.6.11-1+xenial1) ...
    default: Setting up libpython3.6-stdlib:amd64 (3.6.11-1+xenial1) ...
    default: Setting up python3.6 (3.6.11-1+xenial1) ...
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://security.ubuntu.com/ubuntu xenial-security InRelease
    default: Hit:3 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease
    default: Hit:5 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial InRelease
    default: Hit:6 http://archive.ubuntu.com/ubuntu xenial-backports InRelease
    default: Reading package lists...
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: changed: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: [WARNING]: Updating cache and auto-installing missing dependency: python-apt
    default: changed: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Add jonathonf PPA] ************************************************
    default: fatal: [localhost]: FAILED! => {"changed": false, "msg": "Invalid repository string: /etc/apt/sources.list.d/jonathonf-ubuntu-python-3_6-xenial.list"}
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=7    changed=4    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   
The SSH command responded with a non-zero exit status. Vagrant
assumes that this means the command failed. The output for this command
should be in the log above. Please read the output to determine what
went wrong.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant destroy

    default: Are you sure you want to destroy the 'default' VM? [y/N]     default: Are you sure you want to destroy the 'default' VM? [y/N] N
==> default: The VM 'default' will not be destroyed, since the confirmation
==> default: was declined.
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi Vagrantfile 
[sudo] nykim의 암호: 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant provision
==> default: Running provisioner: shell...
    default: Running: inline script
    default: Hit:1 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:2 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Get:4 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Hit:5 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial InRelease
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Fetched 325 kB in 3s (89.7 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: software-properties-common is already the newest version (0.96.20.9).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmpbsax9fv0/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmpbsax9fv0/pubring.gpg' created
    default: gpg: 
    default: requesting key 7BB9C367 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmpbsax9fv0/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 7BB9C367: public key "Launchpad PPA for Ansible, Inc." imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: ansible is already the newest version (2.9.11-1ppa~xenial).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: gpg: 
    default: keyring `/tmp/tmp_6mhkrkg/secring.gpg' created
    default: gpg: 
    default: keyring `/tmp/tmp_6mhkrkg/pubring.gpg' created
    default: gpg: 
    default: requesting key 6A755776 from hkp server keyserver.ubuntu.com
    default: gpg: 
    default: /tmp/tmp_6mhkrkg/trustdb.gpg: trustdb created
    default: gpg: 
    default: key 6A755776: public key "Launchpad PPA for deadsnakes" imported
    default: gpg: 
    default: Total number processed: 1
    default: gpg: 
    default:               imported: 1
    default:   (RSA: 1)
    default: OK
    default: WARNING: 
    default: apt
    default:  
    default: does not have a stable CLI interface. 
    default: Use with caution in scripts.
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: python3.6 is already the newest version (3.6.11-1+xenial1).
    default: 0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
    default: Get:1 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]
    default: Hit:2 http://ppa.launchpad.net/ansible/ansible/ubuntu xenial InRelease
    default: Hit:3 http://ppa.launchpad.net/deadsnakes/ppa/ubuntu xenial InRelease
    default: Hit:4 http://archive.ubuntu.com/ubuntu xenial InRelease
    default: Get:5 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
    default: Get:6 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
    default: Fetched 325 kB in 2s (120 kB/s)
    default: Reading package lists...
    default: Reading package lists...
    default: Building dependency tree...
    default: Reading state information...
    default: E
    default: : 
    default: Unable to locate package python-properties-common
    default: E
    default: : 
    default: Command line option --update is not understood in combination with the other options
    default: E
    default: : 
    default: Command line option --update is not understood in combination with the other options
    default: [WARNING]: provided hosts list is empty, only localhost is available. Note that
    default: the implicit localhost does not match 'all'
    default: 
    default: PLAY [localhost] ***************************************************************
    default: 
    default: TASK [Gathering Facts] *********************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Make app directory] ***********************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Symbolic link] ****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Set timezone to Asia/Seoul] ***************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Update and upgrade apt packages] **********************************
    default: ok: [localhost]
    default: 
    default: TASK [init : fix dependencies] *************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : update again] *****************************************************
    default: ok: [localhost]
    default: 
    default: TASK [init : Install basic packages] *******************************************
    default: changed: [localhost]
    default: 
    default: PLAY RECAP *********************************************************************
    default: localhost                  : ok=9    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ sudo vi Vagrantfile 
nykim@nykim-System-Product-Name:~/바탕화면/vagrant$ ./vagrant ssh
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.4.0-185-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

 * Are you ready for Kubernetes 1.19? It's nearly here! Try RC3 with
   sudo snap install microk8s --channel=1.19/candidate --classic

   https://www.microk8s.io/ has docs and details.

0 packages can be updated.
0 updates are security updates.


*** System restart required ***

This system is built by the Bento project by Chef Software
More information can be found at https://github.com/chef/bento
vagrant@vagrant:~$ python --version
Python 2.7.12
vagrant@vagrant:~$ python3.6
Python 3.6.11 (default, Jun 29 2020, 05:15:03) 
[GCC 5.4.0 20160609] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> 
KeyboardInterrupt
>>> exit
Use exit() or Ctrl-D (i.e. EOF) to exit
>>> exit()
vagrant@vagrant:~$ ansible --version
ansible 2.9.11
  config file = /etc/ansible/ansible.cfg
  configured module search path = [u'/home/vagrant/.ansible/plugins/modules', u'/usr/share/ansible/plugins/modules']
  ansible python module location = /usr/lib/python2.7/dist-packages/ansible
  executable location = /usr/bin/ansible
  python version = 2.7.12 (default, Jul 21 2020, 15:19:50) [GCC 5.4.0 20160609]
vagrant@vagrant:~$ cd /va
vagrant/ var/     
vagrant@vagrant:~$ cd /va
vagrant/ var/     
vagrant@vagrant:~$ cd /va
vagrant/ var/     
vagrant@vagrant:~$ cd /vagrant/
vagrant@vagrant:/vagrant$ ls
ansible  vagrant  Vagrantfile
vagrant@vagrant:/vagrant$ git --version
git version 2.7.4
vagrant@vagrant:/vagrant$ unzip -v
UnZip 6.00 of 20 April 2009, by Debian. Original by Info-ZIP.

Latest sources and executables are at ftp://ftp.info-zip.org/pub/infozip/ ;
see ftp://ftp.info-zip.org/pub/infozip/UnZip.html for other sites.

Compiled with gcc 5.2.1 20151119 for Unix (Linux ELF).

UnZip special compilation options:
        ACORN_FTYPE_NFS
        COPYRIGHT_CLEAN (PKZIP 0.9x unreducing method not supported)
        SET_DIR_ATTRIB
        SYMLINKS (symbolic links supported, if RTL and file system permit)
        TIMESTAMP
        UNIXBACKUP
        USE_EF_UT_TIME
        USE_UNSHRINK (PKZIP/Zip 1.x unshrinking method supported)
        USE_DEFLATE64 (PKZIP 4.x Deflate64(tm) supported)
        UNICODE_SUPPORT [wide-chars, char coding: UTF-8] (handle UTF-8 paths)
        LARGE_FILE_SUPPORT (large files over 2 GiB supported)
        ZIP64_SUPPORT (archives using Zip64 for large files supported)
        USE_BZIP2 (PKZIP 4.6+, using bzip2 lib version 1.0.6, 6-Sept-2010)
        VMS_TEXT_CONV
        WILD_STOP_AT_DIR
        [decryption, version 2.11 of 05 Jan 2007]

UnZip and ZipInfo environment options:
           UNZIP:  [none]
        UNZIPOPT:  [none]
         ZIPINFO:  [none]
      ZIPINFOOPT:  [none]
vagrant@vagrant:/vagrant$ 

```

