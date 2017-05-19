import json

json_string = '{"domain_name": "domain", "dc": "dc", "env": "env1", "node_count": 2}'
parsed_json = json.loads(json_string)

def getListenAddr():
    ''' get the listen address from json'''
    server_list=[]

    for node in range(1, parsed_json['node_count']+1):
        server_list.append(node)
        server = {}
        server['name'] = '{}{}0{}{}_server'.format(parsed_json['domain_name'], parsed_json['dc'], node, parsed_json['env'])
        server['listenAddress'] = '{}{}0{}-{}'.format(parsed_json['domain_name'], parsed_json['dc'], node, parsed_json['env'])
        server['machineName'] = 'nodeMachine0{}'.format(node)
        server['clusterName'] = '{}_cluster1'.format(parsed_json['domain_name'])
        server_list.append(server)

    return server_list

def getMachineName():
    machine_list=[]
    for node in range(1, parsed_json['node_count']+1):
        machine_list.append('node_machine0{}'.format(node))
    return machine_list

def getServerName():
    server_list=[]
    for node in range(1, parsed_json['node_count']+1):
        server_list.append('{}{}0{}{}_server'.format(parsed_json['domain_name'], parsed_json['dc'], node, parsed_json['env']))
    return server_list

def getCluster():
    return ('{}_cluster1'.format(parsed_json['domain_name']))

if __name__ == '__main__':
    print(getListenAddr())
    print(getMachineName())