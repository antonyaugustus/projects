import json

#json_string = '{"first_name": "Guido", "last_name":"Rossum"}'

json_string = '{"domain_name": "mmae", "dc": "idc", "node_count": 2, "env": "env1"}'
parsed_json = json.loads(json_string)

def getHostName():
    for node in range(1, parsed_json['node_count']+1):
        print('{}{}0{}-{}'.format(parsed_json['domain_name'], parsed_json['dc'], node, parsed_json['env'] ))

def getServerName():
    for node in range(1, parsed_json['node_count']+1):
        print('{}{}0{}{}_server'.format(parsed_json['domain_name'], parsed_json['dc'], node, parsed_json['env']))

def getCluster():
    print('{}_cluster1'.format(parsed_json['domain_name']))

if __name__ == '__main__':
    getHostName()
    getServerName()
    getCluster()