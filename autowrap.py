#! /usr/bin/env python
import parser
import sys
from utility import *

def out(s): sys.stdout.write(s)

def levelout(s, n): 
    out(n*"  ")
    out(s)

def newline(): out("\n")

def pretty(pt, depth=0):
    if type(pt) == list:        
        for node in pt:
            pretty(node, depth+1)        
        return            
    if type(pt) == type(""):
        levelout(pt, depth+1)        
        newline()
    else:
        levelout(pt.__name__, depth)
        out(": ")
        out(pt.__name__.line)
        newline()
        pretty(pt.what, depth+1)

def maketree(pt):
    if type(pt) == list:        
        ren = []
        for node in pt:
            ren.append(maketree(node))
        return ren

    if type(pt) == type(""):
        return pt

    else:
        name = pt.__name__
        line = pt.__name__.line        
        return tbl(name)(name, line, maketree(pt.what))

class Node(object):
    def __init__(self, args):
        self.line = args[1]
        self.args = args[2:]

    def cls_name(self):
        return self.__class__.__name__

    def error(self, msg):
        print "wrapper : Error @ %s" % self.line
        print msg        
        raise ValueError("ally: Error @ %s" % self.line)


class symbol(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
    def show(self):
        val = self.args[0]
        if val in ["unsigned", "const"]:
            return ""        
        if val.startswith("sf"):
            val =  val[2:]
        if val.endswith("*"):
            val = "*" + val[:-1]

        return val
        

class ident(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
    def show(self):        
        temp = ''.join(x.show() for x in self.args[0])
        return temp

class const(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
    def show(self):        
        return ""

class unsigned(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
    def show(self):        
        return ""

class void(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
    def show(self):        
        return ""


class idgroup(Node):
    def __init__(self, *args):
        Node.__init__(self, args)

    def cshow(self):
        mapping = {
            "Bool": "bool",
            "sfInt32": "int32"
            }
        ctype = "sf" + self.args[0][0].show().capitalize()
        cname = self.args[0][1].show()

        return '%s(%s)' % (ctype, cname)

    def show(self):        
        temp = ' '.join([x.show() for x in reversed(self.args[0])])
        return temp


cargs = [1]

class parameterlist(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
        
    def uglyhack(self):
        args = []
        for pair in self.args[0][1:]:
            args.append(pair.cshow())

        tmp = ', '.join(args)
        if len(args) > 0:
            tmp = ", " + tmp
            
        cargs[0] = tmp

    def show(self):              
        self.uglyhack()
        result = ", ".join([a.show() for a in self.args[0]][1:]) 
        return result



class proto(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
    def show(self):        
        rtype = self.args[0][0].show()
        if rtype.startswith("sf"):
            rtype = rtype[2:]

        if rtype.endswith("*"):
            rtype = "*" + rtype[0:-1]

        fname = self.args[0][1].show()
        gofname = fname
        if "_" in fname:
            gofname = fname.split("_")[1].capitalize()
        rectype = fname.split("_")[0]
        parms = self.args[0][2].show()
        
        cargsf = str(cargs[0]) 


        s = '''
func (self %s) %s(%s) %s { 
    return C.sf%s(self.Cref%s);
}
            '''
        return s % (rectype, gofname, parms, rtype, fname, cargsf)

class some(Node):
    def __init__(self, *args):
        Node.__init__(self, args)
    def show(self):        
        temp = ''.join(x.show() for x in self.args[0])
        return temp


def tbl(n):
    d = {
        "const":const,
        "unsigned":unsigned,
        "void":void,
        "symbol":symbol,
        "idgroup":idgroup,
        "parameterlist":parameterlist,
        "proto":proto,
        "some":some,
        }
    return d[n]



import sys
if __name__ == "__main__":    
    text = sys.stdin.read()
    for line in text.split("\n"):
        
        if "///////" in line: continue
        line = line.replace("\brief ", "")
        line = line.replace("///", "//")
        if "#include" in line:
            print line
        if line.startswith("///"):
            print line[1:]
            continue
        if line.startswith("//"):
            print line

        # anything after here must be a function to parse
        if "//" in line:
            continue

        if "(" not in line:
            continue

        

        if line.startswith("CSFML_"):
            txt = line.split("_API ")[1]
            print "// " + txt
            tree = parser.go(txt, parser.proto)            
            t = maketree(tree[0])
            print t[0].show()
            
