package main

import (
	"testing"
)

func Test_parseCorefile(t *testing.T) {
	want := `.:53 {
    errors
    health {
        lameduck 5s
    }
    ready
    kubernetes cluster.local in-addr.arpa ip6.arpa {
        pods insecure
        fallthrough in-addr.arpa ip6.arpa
        ttl 30
    }
    prometheus :9153
    forward . /etc/resolv.conf {
        max_concurrent 1000
    }
    cache 30
    loop
    reload
    loadbalance
}
`
	corefile, err := parseCorefile(want)
	if err != nil {
		t.Fatal(err)
	}
	got := corefile.ToString()
	if got != want {
		t.Fatalf("\nwant \n%s \ngot \n%s\n", want, got)
	}
}

func Test_setUpstream(t *testing.T) {
	rawConfig := `.:53 {
        errors
        health {
           lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf {
           max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }`
	rawConfig = `.:53 {
        	errors
        	health {
        		lameduck 5s
        	}
        	ready
        	kubernetes caas.com in-addr.arpa ip6.arpa {
        	   pods insecure
        	   fallthrough in-addr.arpa ip6.arpa
        	   ttl 30
        	}
        	prometheus :9153
        	forward . /etc/resolv.conf
        	cache 10
        	loop
        	reload 10s
        	loadbalance
        }

        :53 {
            errors
            cache 10
        	loadbalance
            hosts {
                1.1.1.1 .caassystem.com
                1.1.1.1 .caassystem.com
                fallthrough
            }
        	forward . 11.11.11.11:5300
        }`
	corefile, err := parseCorefile(rawConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(corefile)
	upstream, err := setUpstream(corefile, "1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(upstream)
}
