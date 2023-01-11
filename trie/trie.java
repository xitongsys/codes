    static class Node {
        long cnt;
        Node[] next;
        public Node(){
            cnt = 0;
            next = new Node[26];
        }
    }

    static void add(String s, Node node, int i){
        while(true) {
            int ns = s.length();
            if (i == ns) {
                node.cnt++;
                return;
            }
            int ci = s.charAt(i) - 'a';
            if (node.next[ci] == null) {
                node.next[ci] = new Node();
            }
            //add(s, node.next[ci], i + 1);
            node = node.next[ci];
            i++;
        }
    }
